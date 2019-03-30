import React, { useState, useEffect } from 'react';
import api from '../api';

import JSONBox from '../components/JSONBox';
import BlogBoxes from '../components/BlogBoxes';

import BlogContext from '../contexts/BlogContext';

import Blog from '../types/blog';

type MyProps = {
	history: any;
	match: any;
	bContext: any;
};

const EachBlogPage = (props: MyProps) => {
	const { bContext } = props;

	const [ blog, setBlog ]: any = useState(bContext.blogs.find(({ _id }: any) => _id === props.match.params.blogId));

	useEffect(
		() => {
			setBlog(bContext.blogs.find(({ _id }: any) => _id === props.match.params.blogId));
		},
		[ bContext.blogs ]
	);

	// Edits Blog, current => current blog object, updates => updated propserties
	const editBlog = async (current: Blog, updates: Blog) => {
		const blogId = current._id;

		delete current._id;
		delete updates._id;

		try {
			const response = await api.put(`/admin/blog/edit/${blogId}`, {
				...current,
				...updates
			});
			bContext.updateBlog(response.data);
		} catch (error) {
			// return error;
			throw error.message;
		}
	};

	return (
		<div className="Page-c-00">
			<div className="Page-Container-00">
				<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
					<JSONBox object={blog} saveFunction={editBlog} />
					<div className="Page-Vertical-Box-Container">
						<BlogBoxes bContext={bContext} blog={blog} setBlog={setBlog} saveFunction={editBlog} />
					</div>
				</div>
			</div>
		</div>
	);
};

export default (props: any) => (
	<BlogContext.Consumer>{(bContext) => <EachBlogPage bContext={bContext} {...props} />}</BlogContext.Consumer>
);
