import React, { Component } from 'react';
import $ from 'jquery';
import './style/ArticleList.css';

import Article from './Article';

class ArticleList extends Component {

    constructor(props) {
        super(props);
        this.state = { articles: [] };
    }

    componentDidMount() {
        const defaultTopic = "finance";
        this.getAPIData(defaultTopic);
    }

    getAPIData(topic) {
        const url = "http://localhost:8080/graphql";
        const query = "{ articles(topic: \"" + topic +"\") { Title Author Confidence Content Description PublishedAt TagName Title Url UrlToImage } }";
        $.post(url, {
            query: query
        },
        (response, status) => {
            if (status === "success") {
                this.setState({ articles: response.data.articles });
            }
        });
    }

    renderArticles() {
        return this.state.articles.map(article =>
            <Article key={article.Title} article={article}/>
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