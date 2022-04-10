package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assetStrings(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		_, got := dictionary.Search("house")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertNoError(t, err)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newValue := "this is a new test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, newValue)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newValue := "new value"
		err := dictionary.Update(key, newValue)
		assertNoError(t, err)
		assertDefinition(t, dictionary, key, newValue)
	})
	t.Run("new key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newKey := "new key"
		newValue := "new value"
		err := dictionary.Update(newKey, newValue)
		assertError(t, err, ErrWordDoesNotExists)
		assertDefinition(t, dictionary, key, value)
	})

}

func TestDelete(t *testing.T) {
	key := "test"
	value := "justa a test"
	dictionary := Dictionary{key: value}

	dictionary.Delete(key)
	_, err := dictionary.Search(key)
	assertError(t, err, ErrNotFound)

}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, value string) {
	got, err := dictionary.Search(key)
	t.Helper()
	assertNoError(t, err)
	assetStrings(t, got, value)
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("should find added word:", err)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assetStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}
