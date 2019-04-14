import { CREATE_BLOG, DELETE_BLOG_BY_ID, EDIT_BLOG_BY_ID, READ_BLOGS, READ_BLOG_BY_ID } from '../actions/action_types';

export default (state = [], action: any) => {
	switch (action.type) {
		case CREATE_BLOG:
			return action.data;
		case DELETE_BLOG_BY_ID:
			return [ ...state.filter(({ _id }: any) => _id !== action.data._id) ];
		case EDIT_BLOG_BY_ID:
			return [ ...state.filter(({ _id }: any) => _id !== action.data._id), action.data ];
		case READ_BLOGS:
			return action.data;
		default:
			return state;
	}
};
