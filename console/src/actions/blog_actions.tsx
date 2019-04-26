import { CREATE_BLOG, DELETE_BLOG_BY_ID, EDIT_BLOG_BY_ID, READ_BLOGS, READ_BLOG_BY_ID } from './action_types';

const blogData = {
	title: 'Title',
	description: 'Description',
	description_link: '',
	thumbnail: '',
	author: 'Ritwik Saha',
	formatted_date: 'January 1, 2025',
	html: '',
	markdown: '',
	doc_type: 'markdown',
	is_public: false,
	is_deleted: false
};

export const createBlog = (extraData: any) => (dispatch: any, getState: any, api: any) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.post('/private/blog/new', { ...blogData, ...extraData });
			dispatch({ type: CREATE_BLOG, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const readBlogs = () => (dispatch: any, getState: any, api: any) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.get('/private/blog/all');
			dispatch({ type: READ_BLOGS, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const editBlog = (blogId: string, editData: any) => (dispatch: any, getState: any, api: any) => {
	delete editData._id; // Deleting _id from editdata (can't update ID)

	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.put(`/private/blog/edit/${blogId}`, editData);
			dispatch({ type: EDIT_BLOG_BY_ID, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};

export const deleteBlog = (blogId: string) => (dispatch: any, getState: any, api: any) => {
	return new Promise(async (resolve, reject) => {
		try {
			const response = await api.delete(`/private/blog/delete/${blogId}`);
			dispatch({ type: DELETE_BLOG_BY_ID, data: response.data });
			resolve(response.data);
		} catch (error) {
			console.log(error);
			reject(error);
		}
	});
};
