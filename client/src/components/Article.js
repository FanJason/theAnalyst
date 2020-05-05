import React, { Component } from 'react';
import './style/Article.css';

class Article extends Component {

    getAuthorName(rawAuthorName) {
        return rawAuthorName.split(" ").slice(0, 2).join(" ");
    }

    getPublishedAtText(rawPublishedAt) {
        return rawPublishedAt.substr(0, rawPublishedAt.indexOf("T"));
    }

    render() {
        const article = this.props.article;
        return (
            <div className="Article row">

                <div className="col-lg-4 mb-4">
                    <div className="view overlay hm-white-slight z-depth-1-half">
                        <img src={article.UrlToImage} className="img-fluid" alt="articleImage"/>
                    </div>
                </div>

                <div className="col-lg-7 ml-xl-4 mb-4">
                    <h4 className="mb-3"><strong>{article.Title}</strong></h4>
                    <p>{article.Description}</p>
                    <p>by <a href="/"><strong>{this.getAuthorName(article.Author)}</strong></a>, {this.getPublishedAtText(article.PublishedAt)}</p>
                    <a href={article.Url} target="#" className="btn btn-primary btn-sm">Read more</a>
                </div>

                <div className="col-lg-7 ml-xl-4 mb-4">
                    <h5>Sentiment Analysis: {article.TagName}, Confidence: {article.Confidence}</h5>
                </div>
                
            </div>
        );
    }
}

export default Article;