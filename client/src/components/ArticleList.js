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
        this.getAPIData(this.props.topic);
    }

    getAPIData(topic) {
        const url = "http://localhost:8080/graphql";
        const query = "{ articles(topic: \"" + topic +"\") { Title Author Confidence Content Description PublishedAt TagName Title Url UrlToImage } }";
        $.post(url, {
            query: query
        },
        (response, status) => {
            if (status === "success") {
                if (!!response && !!response.data && !!response.data.articles) {
                    this.setState({ articles: response.data.articles });
                }
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