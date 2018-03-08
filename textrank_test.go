package textrank

import (
	"testing"
	"time"

	"github.com/DavidBelicza/TextRank/rank"
	"github.com/stretchr/testify/assert"
)

func TestOnSingleThread(t *testing.T) {
	rawText := "Over the past fortnight we asked you to nominate your top extensions for the GNOME desktop. And you did just that. Having now sifted through the hundreds of entries, we’re ready to reveal your favourite GNOME Shell extensions. GNOME 3 (which is more commonly used with the GNOME Shell) has an extension framework that lets developers (and users) extend, build on, and shape how the desktop looks, acts and functions. Dash to Dock takes the GNOME Dash — this is the ‘favourites bar’ that appears on the left-hand side of the screen in the Activities overlay — and transforms it into a desktop dock. And just like Plank, Docky or AWN you can add app launchers, rearrange them, and use them to minimise, restore and switch between app windows. Dash to Dock has many of the common “Dock” features you’d expect, including autohide and intellihide, a fixed-width mode, adjustable icon size, and custom themes. My biggest pet peeve with GNOME Shell is its legacy app tray that hides in the bottom left of the screen. All extraneous non-system applets, indicators and tray icons hide down here. This makes it a little harder to use applications that rely on a system tray presence, like Skype, Franz, Telegram, and Dropbox. TopIcons Plus is the quick way to put GNOME system tray icons back where they belong: on show and in reach. The extension moves legacy tray icons from the bottom left of Gnome Shell to the right-hand side of the top panel. A well-stocked settings panel lets you adjust icon opacity, color, padding, size and tray position. Dive into the settings to adjust the sizing, styling and positioning of icons. Like the popular daily stimulant of choice, the Caffeine GNOME extension keeps your computer awake. It couldn’t be simpler to use: just click the empty mug icon. An empty cup means you’re using normal auto suspend rules – e.g., a screensaver – while a freshly brewed cup of coffee means auto suspend and screensaver are turned off. The Caffeine GNOME extension supports GNOME Shell 3.4 or later. Familiar with applications like Guake and Tilda? If so, you’ll instantly see the appeal of the (superbly named) Drop Down Terminal GNOME extension. When installed just tap the key above the tab key (though it can be changed to almost any key you wish) to get instant access to the command line. Want to speed up using workspaces? This simple tool lets you do just that. Once installed you can quickly switch between workspaces by scrolling over the top panel - no need to enter the Activities Overlay!"

	tr := NewTextRank()
	rule := NewDefaultRule()
	language := NewDefaultLanguage()
	algorithmDef := NewDefaultAlgorithm()

	tr.Populate(rawText, language, rule)
	tr.Ranking(algorithmDef)

	assertTheGnomeTestTextDefault(t, tr)

	algorithmChain := NewChainAlgorithm()
	tr.Ranking(algorithmChain)

	assertTheGnomeTestTextChain(t, tr)
}

func TestOnMultiThread(t *testing.T) {
	exitTest := false
	stream := make(chan string)
	tr := NewTextRank()

	go func(tr *TextRank) {
		i := 1

		for {
			rawText := <-stream

			rule := NewDefaultRule()
			language := NewDefaultLanguage()
			algorithm := NewDefaultAlgorithm()

			tr.Populate(rawText, language, rule)
			tr.Ranking(algorithm)

			if i == 5 {
				exitTest = true
			}

			i++
		}
	}(tr)

	go func() {
		rawTexts := []string{
			"Over the past fortnight we asked you to nominate your top extensions for the GNOME desktop.",
			"And you did just that. Having now sifted through the hundreds of entries, we’re ready to reveal your favourite GNOME Shell extensions. GNOME 3 (which is more commonly used with the GNOME Shell) has an extension framework that lets developers (and users) extend, build on, and shape how the desktop looks, acts and functions.",
			"Dash to Dock takes the GNOME Dash — this is the ‘favourites bar’ that appears on the left-hand side of the screen in the Activities overlay — and transforms it into a desktop dock. And just like Plank, Docky or AWN you can add app launchers, rearrange them, and use them to minimise, restore and switch between app windows. Dash to Dock has many of the common “Dock” features you’d expect, including autohide and intellihide, a fixed-width mode, adjustable icon size, and custom themes.",
			"My biggest pet peeve with GNOME Shell is its legacy app tray that hides in the bottom left of the screen. All extraneous non-system applets, indicators and tray icons hide down here. This makes it a little harder to use applications that rely on a system tray presence, like Skype, Franz, Telegram, and Dropbox. TopIcons Plus is the quick way to put GNOME system tray icons back where they belong: on show and in reach. The extension moves legacy tray icons from the bottom left of Gnome Shell to the right-hand side of the top panel. A well-stocked settings panel lets you adjust icon opacity, color, padding, size and tray position. Dive into the settings to adjust the sizing, styling and positioning of icons. Like the popular daily stimulant of choice, the Caffeine GNOME extension keeps your computer awake.",
			"It couldn’t be simpler to use: just click the empty mug icon. An empty cup means you’re using normal auto suspend rules – e.g., a screensaver – while a freshly brewed cup of coffee means auto suspend and screensaver are turned off. The Caffeine GNOME extension supports GNOME Shell 3.4 or later. Familiar with applications like Guake and Tilda? If so, you’ll instantly see the appeal of the (superbly named) Drop Down Terminal GNOME extension. When installed just tap the key above the tab key (though it can be changed to almost any key you wish) to get instant access to the command line. Want to speed up using workspaces? This simple tool lets you do just that. Once installed you can quickly switch between workspaces by scrolling over the top panel - no need to enter the Activities Overlay!",
		}

		for _, rawText := range rawTexts {
			stream <- rawText
		}
	}()

	for !exitTest {
		time.Sleep(time.Second * 1)
	}

	assertTheGnomeTestTextDefault(t, tr)
}

func assertTheGnomeTestTextDefault(t *testing.T, textRank *TextRank) {
	mostPopulars := []string{
		"gnome shell",
		"icons tray",
		"extension gnome",
		"gnome caffeine",
		"key tab",
		"key changed",
		"overlay activities",
		"auto suspend",
		"dock dash",
	}

	phrases := FindPhrases(textRank)
	max := len(mostPopulars) - 1

	for i := 0; i < max; i++ {
		found := func(ph rank.Phrase) bool {
			for _, popular := range mostPopulars {
				expression := ph.Left + " " + ph.Right

				if expression == popular {
					return true
				}
			}

			return false
		}(phrases[i])

		assert.Equal(t, true, found)
	}

	rankForCheck := textRank.GetRankData()

	assert.Equal(t, float32(1), phrases[0].Weight)
	assert.Equal(t, 5, phrases[0].Qty)
	assert.Equal(t, "gnome", phrases[0].Left)
	assert.Equal(t, "shell", phrases[0].Right)
	assert.Equal(t, phrases[0].LeftID, rankForCheck.WordValID[phrases[0].Left])
	assert.Equal(t, phrases[0].RightID, rankForCheck.WordValID[phrases[0].Right])

	singleWords := FindSingleWords(textRank)

	assert.Equal(t, "gnome", singleWords[0].Word)
	assert.Equal(t, float32(1), singleWords[0].Weight)
	assert.Equal(t, 12, singleWords[0].Qty)
	assert.Equal(t, singleWords[0].ID, rankForCheck.WordValID[singleWords[0].Word])

	sentencesByQtyWeight := FindSentencesByWordQtyWeight(textRank, 6)

	assert.Equal(t, 6, len(sentencesByQtyWeight))
	assert.Equal(t, 0, sentencesByQtyWeight[0].ID)
	assert.Equal(t, 2, sentencesByQtyWeight[1].ID)
	assert.Equal(t, 3, sentencesByQtyWeight[2].ID)
	assert.Equal(t, 4, sentencesByQtyWeight[3].ID)
	assert.Equal(t, 7, sentencesByQtyWeight[4].ID)
	assert.Equal(t, sentencesByQtyWeight[4].Value, rankForCheck.SentenceMap[sentencesByQtyWeight[4].ID])

	sentencesByRelWeight := FindSentencesByRelationWeight(textRank, 6)

	assert.Equal(t, 6, len(sentencesByRelWeight))
	assert.Equal(t, 2, sentencesByRelWeight[0].ID)
	assert.Equal(t, 3, sentencesByRelWeight[1].ID)
	assert.Equal(t, 7, sentencesByRelWeight[2].ID)
	assert.Equal(t, 11, sentencesByRelWeight[3].ID)
	assert.Equal(t, 19, sentencesByRelWeight[4].ID)
	assert.Equal(t, sentencesByRelWeight[4].Value, rankForCheck.SentenceMap[sentencesByRelWeight[4].ID])

	sentencesByPhrase := FindSentencesByPhraseChain(textRank, []string{
		"gnome",
		"shell",
		"extension",
	})

	assert.Equal(t, 3, sentencesByPhrase[0].ID)
	assert.Equal(t, 19, sentencesByPhrase[1].ID)
	assert.Equal(t, sentencesByPhrase[1].Value, rankForCheck.SentenceMap[sentencesByPhrase[1].ID])

	sentenceIDStart := 10
	foundSentences := FindSentencesFrom(textRank, sentenceIDStart, 3)

	assert.Equal(t, sentenceIDStart, foundSentences[0].ID)
	assert.Equal(t, sentenceIDStart+1, foundSentences[1].ID)
	assert.Equal(t, sentenceIDStart+2, foundSentences[2].ID)
	assert.Equal(t, 3, len(foundSentences))
	assert.Equal(t, foundSentences[0].Value, rankForCheck.SentenceMap[foundSentences[0].ID])
}

func assertTheGnomeTestTextChain(t *testing.T, textRank *TextRank) {
	mostPopulars := []string{
		"gnome shell",
		"extension gnome",
		"icons tray",
		"gnome caffeine",
		"key tab",
		"key changed",
		"overlay activities",
		"auto suspend",
		"dock dash",
	}

	phrases := FindPhrases(textRank)
	max := len(mostPopulars) - 1

	for i := 0; i < max; i++ {
		found := func(ph rank.Phrase) bool {
			for _, popular := range mostPopulars {
				expression := ph.Left + " " + ph.Right

				if expression == popular {
					return true
				}
			}

			return false
		}(phrases[i])

		assert.Equal(t, true, found)
	}
}
