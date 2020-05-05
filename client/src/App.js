import React, { Component, createRef }from 'react';
import ArticleList from './components/ArticleList';
import Navbar from './components/Navbar';

class App extends Component {

  constructor(props) {
    super(props);
    this.state = { topic: "finance" };
    this.articleList = createRef();
    this.searchByTopic = this.searchByTopic.bind(this);
  }

  searchByTopic(topic) {
    this.setState({ topic: topic }, () => {
      this.articleList.current.getAPIData(topic);
    });
  }

  render() {
    return (
      <div className="App">
        <Navbar searchByTopic={this.searchByTopic}/>
        <ArticleList ref={this.articleList} topic={this.state.topic}/>
      </div>
    );
  }
}

export default App;