import React, { Component }from 'react';
import ArticleList from './components/ArticleList';
import Navbar from './components/Navbar';

class App extends Component {
  render() {
    return (
      <>
      <Navbar/>
      <ArticleList/>
      </>
    );
  }
}

export default App