import React, { useState, useEffect } from 'react';
import api from '../api';

import data, { Data } from '../data/data';
import ContentBoxItem from './ContentBoxItem';

import Project from '../types/project';
import Blog from '../types/blog';

type MyProps = {
	projects: Array<Project>;
	readProjects: (a: any) => void;
};

const ContentBox = (props: MyProps) => {
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
			setObjects(response.data);
		} catch (error) {
			setObjects(false);
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
			<p className="ContentBox-Types-Btn-Container-01">
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
			</p>
			{objects ? (
				<div className="ContentBox-Item-Container-01">
					{objects.map((object: any, i: number) => (
						<React.Fragment key={i}>
							<ContentBoxItem
								key={i}
								object={object}
								objectLink={`/${activeType === 0 ? 'project' : 'blog'}/${object._id}`}
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

export default ContentBox;
