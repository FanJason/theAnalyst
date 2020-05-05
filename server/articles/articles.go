package articles

import (
	"fmt"
	"encoding/json"

	// "github.com/FanJason/theAnalyst/server/env"
	// "github.com/FanJason/theAnalyst/server/request"
	"github.com/FanJason/theAnalyst/server/sentiment"
)

type Response struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Payload struct {
	Data         []string `json:"data"`
}

type Article struct {
	Author       string
	Title        string
	Description  string
	Url          string
	UrlToImage   string
	PublishedAt  string
	Content      string
	TagName      string
	Confidence   float64
}

func getAPIArticles(topic string) []Article{
	// var apiKey = env.GetEnvVariable("NEWS")
	// url := "http://newsapi.org/v2/everything?q=" + topic + "&sortBy=popularity&apiKey=" + apiKey
	// var data Response
	// request.Get(url, &data)
	// return data.Articles
	article1 := Article{
        Author: "Beth Skwarecki on Vitals, shared by Beth Skwarecki to Lifehacker",
        Confidence: 0.94,
        Content: "Welcome back to our weekly discussion thread, now named the Coronavirus Roundtable. Once again Id like to hear what youve been hearing about our current pandemic, especially if it makes you go huh, could that be true?\r\nLast week we talked about the controvers… [+2343 chars]",
        Description: "Welcome back to our weekly discussion thread, now named the Coronavirus Roundtable. Once again I’d like to hear what you’ve been hearing about our current pandemic, especially if it makes you go “huh, could that be true?” Read more...",
        PublishedAt: "2020-04-06T20:45:00Z",
        TagName: "Positive",
        Title: "Cloth Masks, Wet Markets, and 5G Towers Top Our Coronavirus Concerns This Week",
        Url: "https://vitals.lifehacker.com/cloth-masks-wet-markets-and-5g-towers-top-our-coronav-1842710361",
        UrlToImage: "https://i.kinja-img.com/gawker-media/image/upload/c_fill,f_auto,fl_progressive,g_center,h_675,pg_1,q_80,w_1200/rlg6zvdtp2zn9qdmzt38.jpg",
	}
	article2 := Article{
        Author: "Alex Wilhelm",
        Confidence: 0.94,
        Content: "Hello and welcome back to our regular morning look at private companies, public markets and the gray space in between.\r\nA few weeks back we dug into the boom that savings and investing apps and services were enjoying. Companies like Acorns, M1 Finance, Robinh… [+968 chars]",
        Description: "Have you paid more or less attention to your banking and investing-related apps and, more broadly, your financial life?",
        PublishedAt: "2020-05-04T15:46:38Z",
        TagName: "Positive",
        Title: "A turbulent stock market is a boon to investing-focused fintech startups",
        Url: "http://techcrunch.com/2020/05/04/a-turbulent-stock-market-is-a-boon-to-investing-focused-fintech-startups/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2019/10/GettyImages-1144157724.jpg?w=566",
	}
	article3 := Article{
        Author: "Steve O'Hear",
        Confidence: 0.94,
        Content: "Bó, the digital bank developed by RBS-owned Natwest, is to shutter, just 6 months after launching publicly.\r\nThe incumbent bank’s consumer challenger brand was an attempt to build a startup within a larger bank and in the longer term compete with trendy upsta… [+1225 chars]",
        Description: "Bó, the digital bank developed by RBS-owned Natwest, is to shutter, just 6 months after launching publicly. The incumbent bank’s consumer challenger brand was an attempt to build a startup within a larger bank and in the longer term compete with trendy upstar…",
        PublishedAt: "2020-05-01T11:35:03Z",
        TagName: "Positive",
        Title: "Bó, the digital bank developed by RBS-owned Natwest, is to shutter",
        Url: "http://techcrunch.com/2020/05/01/bo-shutter/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2020/05/Bó-card.png?w=400",
	}
	article4 := Article{
        Author: "Romain Dillet",
        Confidence: 0.94,
        Content: "VC firm Partech has raised a new fund focused on seed investments. Named Partech Entrepreneur III, it is the third seed fund from the VC firm. Partech announced the final closing of its previous seed fund in December 2016.\r\nThe firm is looking for companies a… [+1427 chars]",
        Description: "VC firm Partech has raised a new fund focused on seed investments. Named Partech Entrepreneur III, it is the third seed fund from the VC firm. Partech announced the final closing of its previous seed fund in December 2016. The firm is looking for companies at…",
        PublishedAt: "2020-04-28T07:00:12Z",
        TagName: "Positive",
        Title: "Partech raises $100 million seed fund",
        Url: "http://techcrunch.com/2020/04/28/partech-raises-100-million-seed-fund/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2020/04/image005.jpg?w=764",
	}
	article5 := Article{
        Author: "Jake Bright",
        Confidence: 0.94,
        Content: "The economic effects of COVID-19 could delay Africa’s next big IPO that of Nigerian fintech unicorn Interswitch.\r\nIf so, it wouldn’t be the first time the Lagos-based payments company’s plans for going public were postponed; the tech world has been anticipati… [+1573 chars]",
        Description: "The economic effects of COVID-19 could delay Africa’s next big IPO — that of Nigerian fintech unicorn Interswitch. If so, it wouldn’t be the first time the Lagos-based payments company’s plans for going public were postponed; the tech world has been anticipat…",
        PublishedAt: "2020-05-04T18:03:10Z",
        TagName: "Positive",
        Title: "Why COVID-19 could delay Interswitch, Africa’s next big IPO",
        Url: "http://techcrunch.com/2020/05/04/why-covid-19-could-delay-interswitch-africas-next-big-ipo/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2020/04/IPO-africa-1.png?w=753",
	}
	article6 := Article{
        Author: "Manish Singh",
        Confidence: 0.94,
        Content: "Weeks after Facebook invested $5.7 billion in Indian telecom giant Jio Platforms, private equity firm Silver Lake is following suit. \r\nSilver Lake announced on Monday it will be investing 56.56 billion Indian rupees (about $746.74 million) in the top telecom … [+1618 chars]",
        Description: "Weeks after Facebook invested $5.7 billion in Indian telecom giant Jio Platforms, private equity firm Silver Lake is following suit.  Silver Lake announced on Monday it will be investing 56.56 billion Indian rupees (about $746.74 million) in the top telecom o…",
        PublishedAt: "2020-05-04T03:17:50Z",
        TagName: "Positive",
        Title: "Silver Lake to invest $747M in India’s Jio Platforms",
        Url: "http://techcrunch.com/2020/05/03/silver-lake-jio-platforms/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2020/04/GettyImages-1194806874-1.jpg?w=600",
	}
	article7 := Article{
        Author: "Ron Miller",
        Confidence: 0.94,
        Content: "Koch Industries announced today that it has closed on the acquisition of Infor, announced in February. The company never officially announced the purchase price, but sources indicated that it was close to $13 billion, putting it in line to be one of the top 1… [+1574 chars]",
        Description: "Koch Industries announced today that it has closed on the acquisition of Infor, announced in February. The company never officially announced the purchase price, but sources indicated that it was close to $13 billion, putting it in line to be one of the top 1…",
        PublishedAt: "2020-04-06T14:58:53Z",
        TagName: "Positive",
        Title: "Koch Industries closes nearly $13B Infor acquisition",
        Url: "http://techcrunch.com/2020/04/06/koch-industries-closes-nearly-13b-infor-acquisition/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2020/04/GettyImages-96217367.jpg?w=582",
	}
	article8 := Article{
        Author: "Jordan Crook",
        Confidence: 0.94,
        Content: "23andMe. MongoDB. Eventbrite. Evernote. Bird. Square. Tumblr. Unity. YouTube. Xoom.\r\nRoelof Botha has had a board seat in each of these companies, but his list of investments is much, much longer.\r\nThe Sequoia partner and managing director is legendary in Sil… [+1874 chars]",
        Description: "23andMe. MongoDB. Eventbrite. Evernote. Bird. Square . Tumblr. Unity. YouTube. Xoom. Roelof Botha has had a board seat in each of these companies, but his list of investments is much, much longer. The Sequoia partner and managing director is legendary in Sili…",
        PublishedAt: "2020-04-29T14:39:30Z",
        TagName: "Positive",
        Title: "Extra Crunch Live: Join Roelof Botha for a live Q&A on May 6 at 2pm ET/11am PT",
        Url: "http://techcrunch.com/2020/04/29/extra-crunch-live-join-roelof-botha-for-a-live-qa-on-may-6-at-2pm-et-11am-pt/",
        UrlToImage: "https://techcrunch.com/wp-content/uploads/2020/04/Screen-Shot-2020-04-29-at-9.31.40-AM.png?w=609",
	}
	var articles []Article
	articles = append(articles, article1)
	articles = append(articles, article2)
	articles = append(articles, article3)
	articles = append(articles, article4)
	articles = append(articles, article5)
	articles = append(articles, article6)
	articles = append(articles, article7)
	articles = append(articles, article8)
	return articles
}

func getTitles(articles []Article) []string {
	var titles []string
	for j := 0;  j < 8; j++ {
		titles = append(titles, articles[j].Title)
	}
	return titles
}

func getContent(titles []string) string {
	payload := &Payload{
		Data: titles,
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("failed to parse titles: %v", err)
	}
	return string(bytes)
}

func setSentimentFields(currentArticle * Article, tagName string, confidence float64) {
	currentArticle.TagName = tagName
	currentArticle.Confidence = confidence
}

func getArticlesWithSentiment(articles []Article, response []sentiment.Response) []Article {
	var result []Article
	for j := 0;  j < 8; j++ {
		article := articles[j]
		// classification := response[j].Classifications[0]
		classification := response[0].Classifications[0]
		setSentimentFields(&article, classification.Tag_Name, classification.Confidence)
		result = append(result, article)
	}
	return result
}

func GetArticles(topic string) []Article {
	articles := getAPIArticles(topic)
	titles := getTitles(articles)
	response := sentiment.AnalyzeSentiment(getContent(titles))
	results := getArticlesWithSentiment(articles, response)
	return results
}
