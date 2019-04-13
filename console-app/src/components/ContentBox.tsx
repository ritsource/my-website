import React, { useState, useEffect } from 'react';
import api from '../api';
import { MdLibraryAdd } from 'react-icons/md';

import ContentBoxItem from './ContentBoxItem';
import BlogContext from '../contexts/BlogContext';
import ProjectContext from '../contexts/ProjectContext';

import Project from '../types/project';
import Blog from '../types/blog';

type MyProps = {
	projects: Array<Project>;
	readProjects: (a: any) => void;
	pContext: any;
	bContext: any;
};

const ContentBox = (props: MyProps) => {
	const { pContext, bContext } = props;

	const contentTypes = [ 'projects', 'blogs' ];

	const [ objects, setObjects ]: any = useState([]);
	const [ activeType, setActiveType ] = useState(0);

	const fetchData = async () => {
		let reqPath = '';

		if (contentTypes[activeType] === 'projects') {
			reqPath = '/admin/project/all';
		} else if (contentTypes[activeType] === 'blogs') {
			reqPath = '/admin/blog/all';
		}

		try {
			const response = await api.get(reqPath);
			pContext.readProjects(response.data);
			setObjects(response.data);
		} catch (error) {
			setObjects(false);
		}
	};

	// Creates new Project with Default Values
	const CreateNew = async (cType: string) => {
		// cType - Content Type, Proejct or Blog
		try {
			const response = await api.post(
				`/admin/${cType}/new`,
				cType === 'project'
					? { ...projectData, title: `Project ${objects.length + 1}` }
					: { ...blogData, title: `Project ${objects.length + 1}` }
			);
			// if (cType === 'project') {
			// 	pContext.addProject(response.data);
			// } else {
			// 	bContext.addBlog(response.data);
			// }
			// setObjects([ response.data, ...objects ]);
			await fetchData();
			const element = document.querySelector('.ContentBox-Item-Container-01');
			if (element) {
				element.scrollTop = element.scrollHeight;
			}
		} catch (error) {
			// return error;
			throw error.message;
		}
	};

	useEffect(
		() => {
			setObjects([]);
			fetchData();
		},
		[ activeType ]
	);

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
					onClick={() => {
						CreateNew(activeType === 0 ? 'project' : 'blog');
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
								objectLink={`/${activeType === 0 ? 'projects' : 'blogs'}/${object._id}`}
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

export default (props: any) => (
	<ProjectContext.Consumer>
		{(pContext) => (
			<BlogContext.Consumer>
				{(bContext) => <ContentBox bContext={bContext} pContext={pContext} {...props} />}
			</BlogContext.Consumer>
		)}
	</ProjectContext.Consumer>
);

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

const blogData = {
	title: 'Title',
	description: 'Description',
	author: 'Ritwik Saha',
	formatted_date: 'January 1, 2025',
	html: '',
	markdown: '',
	doc_type: 'markdown',
	is_public: false,
	is_deleted: false
};
