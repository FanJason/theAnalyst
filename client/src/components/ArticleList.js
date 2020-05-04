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
        this.getAPIData();
    }

    getAPIData() {
        const url = "http://localhost:8080/graphql";
        const query = "{articles {Title Author Description Content Url UrlToImage PublishedAt TagName Confidence }}";
        $.post(url, {
            query: query
        },
        (response, status) => {
            if (status === "success") {
                this.setState({ articles: response.data.articles });
            }
        });
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

export default ArticleList;