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
        return (
            <div className="Article">
                {this.props.article.Title}
                <br/>
                {this.props.article.TagName}
                <br/>
                Confidence: {this.props.article.Confidence}
                <br/>
                <button onClick={this.removeArticle}>Remove Me!</button>
            </div>
        );
    }
}

export default Article;