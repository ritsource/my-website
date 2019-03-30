import React, { useState, useEffect } from 'react';

import SubBoxName from './SubBoxName';
import SubBoxDoc from './SubBoxDoc';
import SubBoxBool from './SubBoxBool';

import Blog from '../types/blog';

type MyProps = {
	blog: Blog;
	bContext: any;
	setBlog: () => void;
	saveFunction: (c: Blog, u: Blog) => void;
};

const BlogBoxes = (props: MyProps) => {
	const { blog, saveFunction } = props;

	return (
		<div className="ProjectBoxes-c-00">
			<SubBoxName object={blog} saveFunction={saveFunction} />
			<SubBoxDoc object={blog} saveFunction={saveFunction} />
			<SubBoxBool object={blog} saveFunction={saveFunction} />
		</div>
	);
};

export default BlogBoxes;
