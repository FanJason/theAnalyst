import React, { Component } from 'react';
import './style/Article.css';

class Article extends Component {

    constructor(props) {
        super(props);
        this.removeArticle = this.removeArticle.bind(this);
    }

    removeArticle() {
        this.props.removeArticle(this.props.article.Title);
    }

    render() {
        const article = this.props.article;
        const author = article.Author.split(" ").slice(0, 2).join(" ");
        return (
            <div className="Article row">

                <div class="col-lg-4 mb-4">
                    <div class="view overlay hm-white-slight z-depth-1-half">
                    <img src={article.UrlToImage} class="img-fluid" alt="articleImage"/>
                    <a href="/">
                        <div class="mask"></div>
                    </a>
                    </div>
                </div>

                <div class="col-lg-7 ml-xl-4 mb-4">
                    <h4 class="mb-3"><strong>{article.Title}</strong></h4>
                    <p>{article.Description}</p>
                    <p>by <a href="/"><strong>{author}</strong></a>, {article.PublishedAt}</p>
                    <a href={article.Url} target="#" class="btn btn-primary btn-sm">Read more</a>
                </div>

                <div class="col-lg-7 ml-xl-4 mb-4">
                    <h5>Sentiment Analysis: {article.TagName}, Confidence: {article.Confidence}</h5>
                </div>
                
            </div>
        );
    }
}

export default Article;