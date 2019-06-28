package emoji

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/url"
	"path/filepath"
	"slack-emoji-manager/utils"
	"sort"
	"strings"
)

const (
	emojiDir = "emojis"
)

type Emoji struct {
	name string
	url  string
}

type Alias struct {
	name  string
	alias string
}

func GetEmojis(token string, download bool) error {
	api := slack.New(token)
	emojis, err := api.GetEmoji()
	if err != nil {
		return err
	}

	emojiFiles, aliases := splitEmojis(emojis)

	fmt.Printf("%d aliases\n------\n", len(aliases))
	sortedAlias := make([]string, 0)
	for key := range aliases {
		sortedAlias = append(sortedAlias, key)
	}
	sort.Strings(sortedAlias)

	for _, name := range sortedAlias {
		alias := aliases[name]
		fmt.Printf("%s - %s\n", alias.name, alias.alias)
	}

	fmt.Println()

	if download {
		err := utils.CreateOrClearDir(emojiDir)
		if err != nil {
			return err
		}
	}

	fmt.Printf("%d emojis\n------\n", len(emojiFiles))
	sortedEmojis := make([]string, 0)
	for key := range emojiFiles {
		sortedEmojis = append(sortedEmojis, key)
	}
	sort.Strings(sortedEmojis)

	for _, name := range sortedEmojis {
		emoji := emojiFiles[name]
		fmt.Printf("%s - %s\n", emoji.name, emoji.url)
		if download {
			err := emoji.Download(emojiDir)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//func downloadAllEmojis(emojis map[string]string) {
//	var wg sync.WaitGroup
//	for name, uri := range emojis {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			if !isAlias(uri) {
//				err := downloadEmoji(name, uri)
//				if err != nil {
//					log.Fatal(err)
//				}
//			}
//		}()
//	}
//}

func (emoji *Emoji) Download(dir string) error {
	ext, err := emoji.GetFileExtension()
	if err != nil {
		return err
	}
	newFile := fmt.Sprintf("%s%s", emoji.name, ext)
	err = utils.DownloadFile(filepath.Join(dir, newFile), emoji.url)
	if err != nil {
		return err
	}
	return err
}

func (emoji *Emoji) GetFileExtension() (string, error) {
	u, err := url.Parse(emoji.url)
	if err != nil {
		return "", err
	}
	path := u.Path
	splitPath := strings.Split(path, "/")
	filename := splitPath[len(splitPath)-1]
	ext := filepath.Ext(filename)
	return ext, err
}

//func downloadEmoji(name, emojiURL string) error {
//	ext, err := getFileExtension(emojiURL)
//	if err != nil {
//		return err
//	}
//	newFile := fmt.Sprintf("%s%s", name, ext)
//	err = utils.DownloadFile(filepath.Join(emojiDir, newFile), emojiURL)
//	if err != nil {
//		return err
//	}
//	return err
//}

func splitEmojis(emojis map[string]string) (map[string]Emoji, map[string]Alias) {
	aliases := make(map[string]Alias)
	emojiFiles := make(map[string]Emoji)

	for name, uri := range emojis {
		if isAlias(uri) {
			aliases[name] = Alias{
				name:  name,
				alias: strings.TrimPrefix(uri, "alias:"),
			}
		} else {
			emojiFiles[name] = Emoji{
				name: name,
				url:  uri,
			}
		}
	}
	return emojiFiles, aliases
}

//func getFileExtension(emojiURL string) (string, error) {
//	u, err := url.Parse(emojiURL)
//	if err != nil {
//		return "", err
//	}
//	path := u.Path
//	splitPath := strings.Split(path, "/")
//	filename := splitPath[len(splitPath)-1]
//	ext := filepath.Ext(filename)
//	return ext, err
//}

func isAlias(uri string) bool {
	return strings.HasPrefix(uri, "alias")
}
