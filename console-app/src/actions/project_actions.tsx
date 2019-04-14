import {
	CREATE_PROJECT,
	DELETE_PROJECT_BY_ID,
	EDIT_PROJECT_BY_ID,
	READ_PROJECTS,
	READ_PROJECT_BY_ID
} from './action_types';

// Default Content Data
const projectData = {
	title: 'Title',
	description: 'Description',
	link: '',
	html: '',
	markdown: '',
	doc_type: 'markdown',
	is_public: false,
	is_deleted: false
};

export const createProject = (extraData: any) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.post('/admin/project/new', { ...projectData, ...extraData });
			dispatch({ type: CREATE_PROJECT, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const readProjects = () => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.get('/admin/project/all');
			dispatch({ type: READ_PROJECTS, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const editProject = (projectId: string, editData: any) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.put(`/admin/project/edit/${projectId}`, editData);
			dispatch({ type: EDIT_PROJECT_BY_ID, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const deleteProject = (projectId: string) => (dispatch, getState, api) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.delete(`/admin/project/delete/${projectId}`);
			dispatch({ type: DELETE_PROJECT_BY_ID, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};
