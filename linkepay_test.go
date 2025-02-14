package linkepay

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello from mypackage!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
