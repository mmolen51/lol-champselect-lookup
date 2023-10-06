package main
import (
    "fmt"
    "github.com/gocolly/colly"
    "regexp"
    "io/ioutil"
    "encoding/json"
)

type Champion struct { 
    name string     // champ name
    winrate string  // win percent
    total string    // total games
    mastery string  // mastery points
    score string    // total mastery score
}
var bestCounters []Champion

func main() {
    //counters := ugg("graves")
    // counters := []Champion {
    //     { name: "Ivern", winrate: "54.96% WR", total: "655 games", mastery: "", score: "", }, 
    //     { name: "Nunu & Willump", winrate: "54.9% WR", total: "1,550 games", mastery: "", score: "", }, 
    //     { name: "Udyr", winrate: "53.88% WR", total: "1,197 games", mastery: "", score: "", },
    //     { name: "Evelynn", winrate: "53.5% WR", total: "2,529 games", mastery: "", score: "", }, 
    //     { name: "Bel'Veth", winrate: "52.75% WR", total: "2,749 games", mastery: "", score: "", }, 
    //     { name: "Poppy", winrate: "52.69% WR", total: "1,300 games", mastery: "", score: "", }, 
    //     { name:"Nidalee", winrate: "52.62% WR", total: "5,722 games", mastery: "", score: "", }, 
    //     { name: "Fiddlesticks", winrate: "52.46% WR", total: "1,704 games", mastery: "", score: "", }, 
    //     { name: "Talon", winrate: "52.25% WR", total: "911 games", mastery: "", score: "", }, 
    //     { name: "Kindred", winrate: "51.6% WR", total: "1,469 games", mastery: "", score: "", },
    // }
    //printList(counters)

    //championMastery("vex")
        
    //fmt.Println(bestCounters)
    cons := getConfigs()
    championMastery(cons)
    //tryThings("lux")
    //ugg("zed")

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
    fmt.Println(content)
    err = json.Unmarshal(content, &config)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(config)
    return config
}
func tryThings(champ string) {


}

func ugg(champ string) []Champion {
    var role string
    c := colly.NewCollector()
    c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
    fmt.Println("colly ", c) 


    c.OnRequest(func(r *colly.Request) { 
        fmt.Println("Visiting: ", r.URL) 
    }) 
     
    c.OnError(func(_ *colly.Response, err error) { 
        fmt.Println("Something went wrong: ", err) 
    }) 
     
    c.OnResponse(func(r *colly.Response) { 
        fmt.Println("Page visited: ", r.Request.URL) 
    }) 
     
    c.OnHTML(".role-value", func(e *colly.HTMLElement) {
        role = e.ChildText("div") // role

    })
    c.OnHTML(".counters-list .best-win-rate", func(e *colly.HTMLElement) {
        c := Champion{}
        //fmt.Println(e.ChildText(".champion-name")) 
        c.name = e.ChildText(".champion-name")
        //fmt.Println("winrate", e.ChildText(".win-rate"))
        c.winrate = e.ChildText(".win-rate")
        //fmt.Println(e.ChildText(".total-games"))
        c.total = e.ChildText(".total-games")
        bestCounters = append(bestCounters, c)
        //fmt.Println(bestCounters)
    })
     
    c.OnScraped(func(r *colly.Response) { 
        fmt.Println(r.Request.URL, " scraped!")
    })
    url := fmt.Sprintf("https://u.gg/lol/champions/%s/counter", champ)
    err := c.Visit(url)
    //c.Visit("https://scrapeme.live/shop/")

    if err != nil {
        fmt.Printf("failed to visit url: %v\n", err)
        return nil
    }
    fmt.Printf("%s's best counters in %s are\n", champ, role)
    fmt.Println(bestCounters)
    return bestCounters
    //fmt.Println(bestCounters)
}

func championMastery(configs Config) {
    c := colly.NewCollector()
    c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
    fmt.Println("colly ", c) 


    c.OnRequest(func(r *colly.Request) { 
        fmt.Println("Visiting: ", r.URL) 
    }) 
     
    c.OnError(func(_ *colly.Response, err error) { 
        fmt.Println("Something went wrong: ", err) 
    }) 
     
    c.OnResponse(func(r *colly.Response) { 
        fmt.Println("Page visited: ", r.Request.URL) 
    }) 
     
    re := regexp.MustCompile("[0-9]+")
    

    c.OnHTML("tr", func(e *colly.HTMLElement) {
        fmt.Printf("champ is %s ---- ", e.ChildText(".internalLink")) // role
        text := e.ChildText("td")
        nums := re.FindAllString(text, 1)
        if len(nums) > 0 {
            txt := nums[0]
            var level string
            var score string
            if len(txt) > 0 {
                level = txt[0:1]
                score = txt[1:]
            }
            fmt.Printf("nums[0]: %s\n", nums[0])
            fmt.Printf("level: %s, mastery score: %s\n", level, score)
        }
       
        

        // printing all URLs associated with the a links in the page 
        //fmt.Println("%v", e.Attr("href"))

    })
    // c.OnHTML(".counters-list .best-win-rate", func(e *colly.HTMLElement) {
    //     c := Champion{}
    //     //fmt.Println(e.ChildText(".champion-name")) 
    //     c.name = e.ChildText(".champion-name")
    //     //fmt.Println("winrate", e.ChildText(".win-rate"))
    //     c.winrate = e.ChildText(".win-rate")
    //     //fmt.Println(e.ChildText(".total-games"))
    //     c.total = e.ChildText(".total-games")
    //     bestCounters = append(bestCounters, c)
    // })
     
    c.OnScraped(func(r *colly.Response) { 
        fmt.Println(r.Request.URL, " scraped!")
    })
    fmt.Println("cons: ", configs)
    url := fmt.Sprintf("https://championmastery.gg/summoner?summoner=%s&region=NA&lang=en_US", configs.SummonerName)
    err := c.Visit(url)
    //c.Visit("https://scrapeme.live/shop/")

    if err != nil {
        fmt.Printf("failed to visit url: %v\n", err)
        return
      }
    //fmt.Println(champ)
    //fmt.Println(bestCounters)
}

func printList(champList []Champion) {
    var output string
    for _, ch := range champList {
        tab := "\t"
        if len(ch.name) < 8 {
            tab = tab + "\t"
        }
        output = output + fmt.Sprintf("%s%s%s\t%s\n", ch.name, tab, ch.winrate, ch.total)
    }
    fmt.Println(output)
}

