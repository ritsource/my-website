import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import './App.scss';

import Header from './components/Header';
import HomePage from './pages/Home';
import BlogsPage from './pages/Blogs';
import ProjectsPage from './pages/Projects';
import ImagesPage from './pages/Images';
import EachProjectPage from './pages/EachProject';
import NotFoundPage from './pages/NotFound';

class App extends Component {
	render() {
		return (
			<div className="App">
				<BrowserRouter>
					<Header />
					<Switch>
						<Route path="/" exact component={HomePage} />
						<Route path="/blogs" exact component={BlogsPage} />
						<Route path="/projects" exact component={ProjectsPage} />
						<Route path="/images" exact component={ImagesPage} />
						<Route path="/project/:projectId" exact component={EachProjectPage} />
						<Route component={NotFoundPage} />
					</Switch>
				</BrowserRouter>
			</div>
		);
	}
}

export default App;
