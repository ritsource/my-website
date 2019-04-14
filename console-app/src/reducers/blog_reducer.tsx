import { CREATE_BLOG, DELETE_BLOG_BY_ID, EDIT_BLOG_BY_ID, READ_BLOGS, READ_BLOG_BY_ID } from '../actions/action_types';

export default (state, action) => {
	switch (action.type) {
		case CREATE_BLOG:
			return [ ...state, action.data ];
		case DELETE_BLOG_BY_ID:
			return [ ...state.filter(({ _id }) => _id !== action.data._id) ];
		case EDIT_BLOG_BY_ID:
			return [ ...state.filter(({ _id }) => _id !== action.data._id), action.data ];
		case READ_BLOGS:
			return action.data;
		default:
			break;
	}
};
