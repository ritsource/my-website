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
	description_link: '',
	link: '',
	thumbnail: '',
	html: '',
	markdown: '',
	doc_type: 'markdown',
	is_major: true,
	is_public: false,
	is_deleted: false
};

export const createProject = (extraData: any) => (dispatch: any, getState: any, api: any) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.post('/private/project/new', { ...projectData, ...extraData });
			dispatch({ type: CREATE_PROJECT, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const readProjects = () => (dispatch: any, getState: any, api: any) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.get('/private/projects');
			dispatch({ type: READ_PROJECTS, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const editProject = (projectId: string, editData: any) => (dispatch: any, getState: any, api: any) => {
	delete editData._id; // Deleting _id from editdata (can't update ID)

	console.log(editData);

	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.put(`/private/project/edit?id=${projectId}`, editData);
			dispatch({ type: EDIT_PROJECT_BY_ID, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const deleteProject = (projectId: string) => (dispatch: any, getState: any, api: any) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.delete(`/private/project/delete?id=${projectId}`);
			dispatch({ type: DELETE_PROJECT_BY_ID, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};
