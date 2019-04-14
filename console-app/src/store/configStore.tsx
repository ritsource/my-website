import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import api from '../api';

export default () => {
	const store = createStore(combineReducers({}), applyMiddleware(thunk.withExtraArgument(api)));

	return store;
};
