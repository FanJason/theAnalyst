import React, { Component } from 'react';
import $ from 'jquery';
import './style/ArticleList.css';

import Article from './Article';

class ArticleList extends Component {

    constructor(props) {
        super(props);
        this.state = { articles: [] };
        this.removeArticle = this.removeArticle.bind(this);
    }

    componentDidMount() {
        const url = "http://localhost:8080/graphql";
        const query = "{articles {Title Author Description Url UrlToImage PublishedAt TagName Confidence }}";
        $.post(url, {
            query: query
        },
        function(response, status) {
            let newState = this.state
            newState.articles = response.data.articles;
            this.setState(newState);
            console.log(response);
        }.bind(this));
    }

    removeArticle(removeTitle) {
        const filteredArticles = this.state.articles.filter(article => {
            return article.Title !== removeTitle;
        });
        this.setState({ articles: filteredArticles });
    }

    renderArticles() {
        return this.state.articles.map(article =>
            <Article key={article.Title} article={article} removeArticle={this.removeArticle}/>
        );
    }

    render() {
        return (
            <div className="ArticleList">
                {this.renderArticles()}
            </div>
        )
    }
}

export default ArticleList