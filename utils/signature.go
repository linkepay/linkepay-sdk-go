package utils

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/linkepay/linkepay-sdk-go/types"
)

type KeyManager struct {
	privateKey           string
	publicKey            string
	address              string
	payPlatformPublicKey string // pay platform public key
}

func NewKeyManager() *KeyManager {
	return &KeyManager{}
}

func (km *KeyManager) LoadKeys(config *types.Config) {
	km.privateKey = config.PrivateKey
	km.publicKey = config.PublicKey
	km.payPlatformPublicKey = config.PayPlatformPublicKey
}

func (km *KeyManager) SetPrivateKey(privateKey string) {
	km.privateKey = privateKey
}

func (km *KeyManager) SetPublicKey(publicKey string) {
	km.publicKey = publicKey
}

func (km *KeyManager) GenerateKeys() (types.Keys, error) {
	// Generate new ECDSA private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return types.Keys{}, err
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return types.Keys{}, err
	}

	// Convert to hex strings
	km.privateKey = hexutil.Encode(crypto.FromECDSA(privateKey))
	km.publicKey = hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	km.address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return types.Keys{
		PrivateKey: km.privateKey,
		PublicKey:  km.publicKey,
		Address:    km.address,
	}, nil
}

func (km *KeyManager) SaveKeysToFile(privateKeyFile, publicKeyFile string) error {
	if err := os.WriteFile(privateKeyFile, []byte(km.privateKey), 0600); err != nil {
		return fmt.Errorf("failed to save private key: %w", err)
	}
	if err := os.WriteFile(publicKeyFile, []byte(km.publicKey), 0600); err != nil {
		return fmt.Errorf("failed to save public key: %w", err)
	}
	return nil
}

func (km *KeyManager) LoadPrivateKeyFromFile(privateKeyFile, publicKeyFile string) error {
	privateKey, err := os.ReadFile(privateKeyFile)
	if err != nil {
		return fmt.Errorf("failed to read private key file: %w", err)
	}
	km.privateKey = strings.TrimSpace(string(privateKey))

	publicKey, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return fmt.Errorf("failed to read public key file: %w", err)
	}
	km.publicKey = strings.TrimSpace(string(publicKey))

	return nil
}

func (km *KeyManager) SignRequest(path string, body interface{}) (string, error) {
	var message string
	if body != nil {
		sortedBodyStr, err := ToSortedJSON(body)
		if err != nil {
			return "", fmt.Errorf("failed to sort body: %w", err)
		}
		message = path + sortedBodyStr
	} else {
		message = path
	}

	// Hash the message
	messageHash := sha256.Sum256([]byte(message))

	// Parse private key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(km.privateKey, "0x"))
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// Sign the hash
	signature, err := crypto.Sign(messageHash[:], privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %w", err)
	}

	// Convert signature to hex
	signatureHex := hex.EncodeToString(signature)

	// Add 0x prefix
	return "0x" + signatureHex, nil
}

// func (km *KeyManager) VerifySignature(publicKeyHex, message, signatureHex string) (bool, error) {
// 	// Hash the message
// 	messageHash := sha256.Sum256([]byte(message))

// 	// Decode signature from hex
// 	signature, err := hex.DecodeString(strings.TrimPrefix(signatureHex, "0x"))
// 	if err != nil {
// 		return false, fmt.Errorf("failed to decode signature: %w", err)
// 	}

// 	// Get public key bytes
// 	publicKeyBytes, err := hex.DecodeString(strings.TrimPrefix(publicKeyHex, "0x"))
// 	if err != nil {
// 		return false, fmt.Errorf("failed to decode public key: %w", err)
// 	}

// 	// If the public key is 64 bytes, prepend 0x04 to indicate it's uncompressed
// 	if len(publicKeyBytes) == 64 {
// 		publicKeyBytes = append([]byte{0x04}, publicKeyBytes...)
// 	}

// 	// Unmarshal the public key
// 	pubKey, err := crypto.UnmarshalPubkey(publicKeyBytes)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to unmarshal public key: %w", err)
// 	}

// 	// Verify the signature
// 	return crypto.VerifySignature(crypto.FromECDSAPub(pubKey), messageHash[:], signature[:len(signature)-1]), nil
// }

func (km *KeyManager) VerifySignature(publicKey string, signature string, message string) (bool, error) {
	messageHash := sha256.Sum256([]byte(message))

	signature = strings.TrimPrefix(signature, "0x")
	signature = strings.TrimPrefix(signature, "0x")
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false, err
	}
	publicKeyHex := strings.TrimPrefix(publicKey, "0x")
	publicKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return false, err
	}
	// Check if the public key is in the correct format
	if len(publicKeyBytes) != 64 && len(publicKeyBytes) != 65 {
		return false, err
	}

	// If the public key is 64 bytes, prepend 0x04 to indicate it's uncompressed
	if len(publicKeyBytes) == 64 {
		publicKeyBytes = append([]byte{0x04}, publicKeyBytes...)
	}

	// Verify the signature using Ethereum's crypto package
	pubKey, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		return false, err
	}
	if !crypto.VerifySignature(crypto.FromECDSAPub(pubKey), messageHash[:], signatureBytes[:len(signatureBytes)-1]) {
		return false, err
	}
	return true, nil
}

func SignWithPrivateKeyFile(path string, body interface{}, privateKeyFile string) (string, error) {
	km := NewKeyManager()
	if err := km.LoadPrivateKeyFromFile(privateKeyFile, "public_key.txt"); err != nil {
		return "", err
	}
	return km.SignRequest(path, body)
}

func (km *KeyManager) VerifyPlatformSignature(platformPublicKey string, data interface{}, signature string) (bool, error) {
	sortedDataStr, err := ToSortedJSON(data)
	if err != nil {
		return false, errors.New("failed to marshal sorted data")
	}
	ok, err := km.VerifySignature(platformPublicKey, signature, sortedDataStr)
	if err != nil {
		return false, errors.New("failed to verify signature")
	}

	return ok, nil
}

func (km *KeyManager) GetPlatformPublicKey() string {
	return km.payPlatformPublicKey
}

func (km *KeyManager) GetPublicKey() string {
	return km.publicKey
}

func (km *KeyManager) GetPrivateKey() string {
	return km.privateKey
}

func (km *KeyManager) GenerateSortedDataStr(data interface{}) (string, error) {
	sortedDataStr, err := ToSortedJSON(data)
	if err != nil {
		return "", fmt.Errorf("failed to sort data: %w", err)
	}
	return sortedDataStr, nil
}

func (km *KeyManager) GeneratePlatformSignature(sortedDataStr string) (string, error) {
	return km.SignRequest(sortedDataStr, nil)
}
