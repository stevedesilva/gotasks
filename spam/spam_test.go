package spam_test

import (
	"testing"
	"errors"
	"github.com/stevedesilva/learngo/spam"
)

func TestSpammer_should_be_created(t *testing.T) {
	spammer := spam.New()
	if spammer == nil {
		t.Errorf("Expected new Spammer value")
	}
}

func TestSpammer_should_quit_with_error_message(t *testing.T) {

	spammer := spam.New()
	_, err := spammer.Process([]string{"firstarg"})
	if err == nil {
		t.Errorf("Expected Error %v", err)
	}

	if !errors.Is(spam.ErrSpammer,err) {
		t.Errorf("Expected Error %v, got %v.", spam.ErrSpammer, err)
	}
}

func TestSpammer_should_read_args_from_cmdline_without_change(t *testing.T) {
	input := "Here's my spammy page: hTTp://youth-elixir.com"

	spammer := spam.New()

	actual, err := spammer.Process([]string{"firstarg", input})

	if expected := input; expected != actual {
		t.Errorf("Expected %s Actual %s, error %v", expected, actual, err)
	}
}

func TestSpammer_should_read_args_from_cmdline_with_change(t *testing.T) {
	input := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	expected := "Here's my spammy page: http://******************* see you."
	
	spammer := spam.New()
	actual, err := spammer.Process([]string{"firstarg", input})
	
	if expected != actual {
		t.Errorf("Expected %s Actual %s, error %v", expected, actual, err)
	}
}
