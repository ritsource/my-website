import React, { useState, useEffect } from 'react';
import api from '../api';
import Blog from '../types/blog';

const BlogContext = React.createContext({
	blogs: [],
	readBlogs: (a: Array<Blog>) => {},
	addBlog: (p: Blog) => {},
	updateBlog: (p: Blog) => {},
	deleteBlog: (id: string) => {}
});

export const BlogProvider = (props: any) => {
	const [ blogs, setBlogs ] = useState([]);

	useEffect(() => {
		(async () => {
			try {
				const response = await api.get('/admin/blog/all');
				setBlogs(response.data);
			} catch (error) {
				setBlogs([]);
			}
		})();
	}, []);

	const readBlogs = (allBlogs: any) => {
		setBlogs(allBlogs);
	};

	const addBlog = (newBlog: any) => {
		setBlogs(blogs.concat(newBlog));
	};

	const updateBlog = (newBlog: any) => {
		setBlogs(blogs.filter(({ _id }) => _id !== newBlog._id).concat(newBlog));
	};

	const deleteBlog = (blogId: string) => {
		setBlogs([ ...blogs.filter(({ _id }) => _id !== blogId) ]);
	};

	return (
		<BlogContext.Provider
			value={{
				blogs,
				readBlogs,
				addBlog,
				updateBlog,
				deleteBlog
			}}
		>
			{props.children}
		</BlogContext.Provider>
	);
};

export default BlogContext;
