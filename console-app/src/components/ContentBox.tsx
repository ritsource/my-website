import React, { useState, useEffect } from 'react';

import data, { Data } from '../data/data';
import ContentBoxItem from './ContentBoxItem';

const ContentBox = () => {
	const contentTypes = [ 'projects', 'blogs', 'images' ];

	const [ objects, setObjects ]: any = useState([]);
	const [ activeType, setActiveType ] = useState(0);

	const fetchData = () => {
		setObjects(data);
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
					<React.Fragment>
						<span
							onClick={() => setActiveType(index)}
							className={`${index === activeType && 'ContentBox-Active-Span'}`}
						>
							{type}
						</span>{' '}
					</React.Fragment>
				))}
			</p>
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
