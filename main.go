package main
import (
    "fmt"
    "github.com/gocolly/colly"
    "regexp"
    "io/ioutil"
    "encoding/json"
    "os"
    "strings"
)

type Champion struct { 
    name string     // champ name
    winrate string  // win percent
    total string    // total games
    mastery string  // mastery level
    score string    // total mastery score
}
var bestCounters []Champion
var champMastList []Champion
var combined []Champion

func main() {
    args := os.Args
    cname := args[1]
    //fmt.Println(cname)
    counters, role := ugg(cname)
    cons := getConfigs()
    ml := championMastery(cons)
    if len(ml) == 0 {
        fmt.Println("\nI didn't find your summoner name. Please re-enter it with\nlol ign <summoner name>")
        return
    }
    list := combineLists(counters, ml)
    if (role != "") {
        fmt.Printf("\n%s's best counters in %s are\n", strings.Title(cname), role)
    } else {
        fmt.Println("\nI didn't find anything about that champion, please check the spelling.\nDon't include spaces or apostrophes.\n\nTo use, type\nlol <champ name>\n\nexamples:\nlol lux\nlol kaisa\nlol nunu\n")
    }
    printList(list)
}

type Config struct {
    SummonerName string
}

func getConfigs() Config {
    config := Config{}
    content, err := ioutil.ReadFile("configs.json")
	if err != nil {
		fmt.Println(err)
	}
    err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println(err)
	}
    return config
}


func ugg(champ string) ([]Champion, string) {
    var role string
    c := colly.NewCollector()
    c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

    c.OnRequest(func(r *colly.Request) { 
        fmt.Println("Visiting: ", r.URL) 
    }) 
     
    c.OnError(func(_ *colly.Response, err error) { 
        fmt.Println("Something went wrong: ", err) 
    }) 
     
    c.OnResponse(func(r *colly.Response) { 
        //fmt.Println("Page visited: ", r.Request.URL) 
    }) 
     
    c.OnHTML(".role-value", func(e *colly.HTMLElement) {
        role = e.ChildText("div") // role

    })
    c.OnHTML(".counters-list .best-win-rate", func(e *colly.HTMLElement) {
        c := Champion{}
        c.name = e.ChildText(".champion-name")
        c.winrate = e.ChildText(".win-rate")
        c.total = e.ChildText(".total-games")
        bestCounters = append(bestCounters, c)
    })
     
    c.OnScraped(func(r *colly.Response) { 
        //fmt.Println(r.Request.URL, " scraped!")
    })
    url := fmt.Sprintf("https://u.gg/lol/champions/%s/counter", champ)
    err := c.Visit(url)

    if err != nil {
        fmt.Printf("failed to visit url: %v\n", err)
        return nil, ""
    }
    return bestCounters, role
}

func championMastery(configs Config) []Champion {
    c := colly.NewCollector()
    c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

    c.OnRequest(func(r *colly.Request) { 
        fmt.Println("Visiting: ", r.URL) 
    }) 
     
    c.OnError(func(_ *colly.Response, err error) { 
        fmt.Println("Something went wrong: ", err) 
    }) 
     
    c.OnResponse(func(r *colly.Response) { 
        //fmt.Println("Page visited: ", r.Request.URL) 
    }) 
     
    re := regexp.MustCompile("[0-9]+")
    cm := Champion{} 

    c.OnHTML("tr", func(e *colly.HTMLElement) {
        //fmt.Printf("champ is %s ---- ", e.ChildText(".internalLink")) // role
        cm.name = e.ChildText(".internalLink")
        text := e.ChildText("td")
        nums := re.FindAllString(text, 1)
        if len(nums) > 0 {
            txt := nums[0]
            var level string
            var score string
            if len(txt) > 0 {
                level = txt[0:1]
                cm.mastery = level
                score = txt[1:]
                cm.score = score
                champMastList = append(champMastList, cm)
            }
        }
    })
     
    c.OnScraped(func(r *colly.Response) { 
        //fmt.Println(r.Request.URL, " scraped!")
    })
    sn := strings.Replace(configs.SummonerName, " ", "+", -1)
    url := fmt.Sprintf("https://championmastery.gg/summoner?summoner=%s&region=NA&lang=en_US", sn)
    err := c.Visit(url)

    if err != nil {
        fmt.Printf("failed to visit url: %v\n", err)
        return nil
      }
    return champMastList
}

func combineLists(c []Champion, ml []Champion) []Champion {
    for _, ch := range c {
        found := false
        for _, m := range ml {
            if m.name == ch.name {
                nc := Champion{}
                nc.name = ch.name
                nc.winrate = ch.winrate
                nc.total = ch.total
                nc.mastery = m.mastery
                nc.score = m.score
                combined = append(combined, nc)
                found = true
                break
            }
        }
        if (found == false) {
            nc := Champion{}
            nc.name = ch.name
            nc.winrate = ch.winrate
            nc.total = ch.total
            nc.mastery = "0"
            nc.score = "0"
            combined = append(combined, nc)
        }
    }
    return combined
}

func printList(champList []Champion) {
    var output string
    for _, ch := range champList {
        tab := "\t"
        tab2 := "\t"
        if len(ch.name) < 8 {
            tab = tab + "\t"
        }
        if len(ch.total) < 10 {
            tab2 = tab2 + "\t"
        }
        output = output + fmt.Sprintf("%s%s%s in %s%s mastery %s with %s pts\n", ch.name, tab, ch.winrate, ch.total, tab2, ch.mastery, ch.score)
    }
    fmt.Println(output)
}

