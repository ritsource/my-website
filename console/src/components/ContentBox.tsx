import React, { useState, useEffect } from 'react';
import { MdLibraryAdd } from 'react-icons/md';
import { connect } from 'react-redux';

import { createProject, readProjects } from '../actions/project_actions';
import { createBlog, readBlogs } from '../actions/blog_actions';
import ContentBoxItem from './ContentBoxItem';

import Project from '../types/project';
import Blog from '../types/blog';

const ContentBox = (props: any) => {
	const { projects, blogs } = props;

	const contentTypes = [ 'projects', 'blogs' ];

	const [ activeType, setActiveType ] = useState(0);

	const isProject = contentTypes[activeType] === 'projects';

	useEffect(
		() => {
			if (isProject) {
				props.readProjects();
			} else {
				props.readBlogs();
			}
		},
		[ activeType ]
	);

	const objects = isProject ? projects : blogs;

	return (
		<div className="SearchBox-c-00 ContentBox-c-00 Theme-Box-Shadow">
			<h4>Latest Content</h4>
			<p className="SearchBox-Description-P-01">All the latest Projects, Photos, Blogs.</p>
			<p className="ContentBox-Types-Btn-Container-01 Flex-Row-Space-Between">
				<span>
					{contentTypes.map((type, index) => (
						<React.Fragment key={index}>
							<span
								key={index}
								onClick={() => setActiveType(index)}
								className={`${index === activeType && 'ContentBox-Active-Span'}`}
							>
								{type}
							</span>{' '}
						</React.Fragment>
					))}
				</span>

				<button
					className="ContentBox-Add-Btn-02 Flex-Row-Center"
					onClick={async () => {
						if (isProject) {
							await props.createProject({ title: `Project ${projects.length + 1}` });
						} else {
							await props.createBlog({ title: `Blog ${blogs.length + 1}` });
						}
						const element = document.querySelector('.ContentBox-Item-Container-01');
						if (element) {
							element.scrollTop = element.scrollHeight;
						}
					}}
				>
					<MdLibraryAdd style={{ fontSize: '16px', marginRight: '5px' }} />
					{activeType === 0 ? 'New Project' : 'New Blog'}
				</button>
			</p>

			{objects ? (
				<div className="ContentBox-Item-Container-01">
					{objects.map((object: any, i: number) => (
						<React.Fragment key={i}>
							<ContentBoxItem
								key={i}
								object={object}
								objectLink={`${activeType === 0 ? '/projects' : '/blogs'}/${object._id}`}
							/>
						</React.Fragment>
					))}
				</div>
			) : (
				<p style={{ marginTop: '5px' }} className="SearchBox-Description-P-01">
					Network Error, Couldn't fetch Data!
				</p>
			)}
		</div>
	);
};

const mapStateToProps = ({ blogs, projects }: { blogs: Array<Blog>; projects: Array<Project> }) => ({
	projects: projects || [],
	blogs: blogs || []
});

const mapDispatchToProps = (dispatch: (x: any) => void) => ({
	createProject: (x: any) => dispatch(createProject(x)),
	readProjects: () => dispatch(readProjects()),
	createBlog: (x: any) => dispatch(createBlog(x)),
	readBlogs: () => dispatch(readBlogs())
});

export default connect(mapStateToProps, mapDispatchToProps)(ContentBox);
