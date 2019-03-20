import React, { useState, useEffect } from 'react';

import ContentBoxItem from './ContentBoxItem';

const ContentBox = () => {
	const [ objects, setObjects ]: any = useState([]);

	const fetchData = () => {
		setObjects(data);
	};

	useEffect(() => {
		fetchData();
	}, []);

	return (
		<div className="SearchBox-c-00 ContentBox-c-00 Theme-Box-Shadow">
			<h4>Latest Content</h4>
			<p className="SearchBox-Description-P-01">All the latest Projects, Photos, Blogs.</p>
			<div className="ContentBox-Item-Container-01">
				{objects.map((object: Data, i: number) => (
					<React.Fragment>
						<ContentBoxItem key={i} object={object} />
					</React.Fragment>
				))}
			</div>
		</div>
	);
};

export default ContentBox;

export type Data = {
	title: string;
	type: string;
	emoji: string;
	description: string;
	html: string;
	markdown: string;
	link: string;
	imageUrl: string;
};

const data: Array<Data> = [
	{
		title: 'My Life til 19',
		type: 'blog',
		emoji: '🍏',
		description: 'description',
		html: 'http://content.ritwiksaha.com/xyz',
		markdown: 'http://content.ritwiksaha.com/xyz',
		link: 'http://content.ritwiksaha.com/xyz',
		imageUrl: 'http://content.ritwiksaha.com/xyz'
	},
	{
		title: 'Scheduler',
		type: 'project',
		emoji: '🍌',
		description: 'description',
		html: 'http://content.ritwiksaha.com/xyz',
		markdown: 'http://content.ritwiksaha.com/xyz',
		link: 'http://content.ritwiksaha.com/xyz',
		imageUrl: 'http://content.ritwiksaha.com/xyz'
	},
	{
		title: 'Raspi',
		type: 'project',
		emoji: '🍓',
		description: 'description',
		html: 'http://content.ritwiksaha.com/xyz',
		markdown: 'http://content.ritwiksaha.com/xyz',
		link: 'http://content.ritwiksaha.com/xyz',
		imageUrl: 'http://content.ritwiksaha.com/xyz'
	}
];
