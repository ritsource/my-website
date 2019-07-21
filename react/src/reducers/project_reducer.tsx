import {
	CREATE_PROJECT,
	DELETE_PROJECT_BY_ID,
	EDIT_PROJECT_BY_ID,
	READ_PROJECTS,
	READ_PROJECT_BY_ID
} from '../actions/action_types';

export default (state = [], action: any) => {
	switch (action.type) {
		case CREATE_PROJECT:
			return action.data;
		case DELETE_PROJECT_BY_ID:
			return [ ...state.filter(({ _id }: any) => _id !== action.data._id) ];
		case EDIT_PROJECT_BY_ID:
			return [ ...state.filter(({ _id }: any) => _id !== action.data._id), action.data ];
		case READ_PROJECTS:
			return action.data;
		default:
			return state;
	}
};
