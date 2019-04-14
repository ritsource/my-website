import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import api from '../api';

import blogReducer from '../reducers/blog_reducer';
import projectReducer from '../reducers/project_reducer';

export default () => {
	const store = createStore(
		combineReducers({
			blogs: blogReducer,
			projects: projectReducer
		}),
		applyMiddleware(thunk.withExtraArgument(api))
	);

	return store;
};
