import React, { useState, useEffect } from 'react';

import SubBoxName from './SubBoxName';
import SubBoxDoc from './SubBoxDoc';
import SubBoxBool from './SubBoxBool';
import SubBoxDelete from './SubBoxDelete';
import SubBoxUpdate from './SubBoxUpdate';

import Blog from '../types/blog';

type MyProps = {
	blog: Blog;
	setBlog: () => void;
	saveFunction: (c: Blog, u: Blog) => void;
};

const BlogBoxes = (props: MyProps) => {
	const { blog, saveFunction } = props;

	return (
		<div className="ProjectBoxes-c-00">
			<SubBoxName object={blog} saveFunction={saveFunction} isProject={false} />
			<SubBoxDoc object={blog} saveFunction={saveFunction} />
			<SubBoxUpdate correctId={blog._id} />
			<SubBoxBool object={blog} saveFunction={saveFunction} />
			<SubBoxDelete isProject={false} correctId={blog._id} />
		</div>
	);
};

export default BlogBoxes;
