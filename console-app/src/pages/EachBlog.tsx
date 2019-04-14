import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';

import { readBlogs, createBlog, editBlog } from '../actions/blog_actions';
import JSONBox from '../components/JSONBox';
import BlogBoxes from '../components/BlogBoxes';
import LoadingPage from './Loading';

import Blog from '../types/blog';

const EachBlogPage = (props: any) => {
	const { blogs } = props;

	const [ blog, setBlog ]: any = useState(blogs.find(({ _id }: any) => _id === props.match.params.blogId));

	useEffect(
		() => {
			setBlog(blogs.find(({ _id }: any) => _id === props.match.params.blogId));
		},
		[ blogs ]
	);

	return (
		<div className="Page-c-00">
			<div className="Page-Container-00">
				<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
					{blog ? (
						<React.Fragment>
							{/* In the function below c = current-blog-data; u = new-blog-data */}
							<JSONBox object={blog} saveFunction={(c, u) => props.editBlog(c._id, { ...c, ...u })} />
							<div className="Page-Vertical-Box-Container">
								<BlogBoxes
									blog={blog}
									setBlog={setBlog}
									saveFunction={(c, u) => props.editBlog(c._id, { ...c, ...u })}
								/>
							</div>
						</React.Fragment>
					) : (
						<LoadingPage />
					)}
				</div>
			</div>
		</div>
	);
};

const mapStateToProps = ({ blogs }: { blogs: Array<Blog> }) => ({
	blogs: blogs || []
});

const mapDispatchToProps = (dispatch: (x: any) => void) => ({
	readBlogs: () => dispatch(readBlogs()),
	editBlog: (x: string, y: any) => dispatch(editBlog(x, y))
});

export default connect(mapStateToProps, mapDispatchToProps)(EachBlogPage);
