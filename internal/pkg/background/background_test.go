package background

import "testing"

func TestBackground(t *testing.T) {
	t.Run("testing name getter", func(t *testing.T) {
		b := GUILD_ARTISAN
		name := b.name
		if name != GUILD_ARTISAN.name {
			t.Errorf("unexpected, got %v", name)
		}
	})

	t.Run("testing description getter", func(t *testing.T) {
		b := FOLK_HERO
		description := b.description
		if description != FOLK_HERO.description {
			t.Errorf("unexpected, got %v", description)
		}
	})
}
