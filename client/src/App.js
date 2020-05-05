import React, { Component }from 'react';
import ArticleList from './components/ArticleList';
import Navbar from './components/Navbar';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Navbar/>
        <ArticleList/>
      </div>
    );
  }
}

export default App