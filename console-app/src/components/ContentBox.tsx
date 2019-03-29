import React, { useState, useEffect } from 'react';
import api from '../api';

import data, { Data } from '../data/data';
import ContentBoxItem from './ContentBoxItem';

import Project from '../types/project';

type MyProps = {
	projects: Array<Project>;
	readProjects: (a: any) => void;
};

const ContentBox = (props: MyProps) => {
	const contentTypes = [ 'projects', 'blogs' ];

	const [ objects, setObjects ]: any = useState([]);
	const [ activeType, setActiveType ] = useState(0);

	const fetchData = async () => {
		if (contentTypes[activeType] === 'projects') {
			try {
				const response = await api.get('/admin/project/all');
				setObjects(response.data);
			} catch (error) {
				setObjects(false);
			}
		} else if (contentTypes[activeType] === 'blogs') {
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
			<div className="ContentBox-Item-Container-01">
				{objects ? (
					<React.Fragment>
						{objects.map((object: Data, i: number) => (
							<React.Fragment key={i}>
								<ContentBoxItem key={i} object={object} />
							</React.Fragment>
						))}
					</React.Fragment>
				) : (
					<h4>No Data Found</h4>
				)}
			</div>
		</div>
	);
};

export default ContentBox;
