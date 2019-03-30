import React from 'react';
import { Link } from 'react-router-dom';
import { MdBlock, MdLockOutline, MdLockOpen } from 'react-icons/md';

import { Data } from '../data/data';

import Project from '../types/project';
import Blog from '../types/project';
import { object } from 'prop-types';

type MyProps = {
	// this could be Project & Blog, but for convicince let it be one
	object: any;
	objectLink: string;
};

const ContentBoxItem: React.SFC<MyProps> = (props) => {
	// const { title, emoji, description, html, markdown, link, imageUrl } = props.object;
	const { object, objectLink } = props;

	return (
		<div className="ContentBoxItem-c-00 Flex-Row-Space-Between">
			<p>
				<Link to={objectLink}>{object.title}</Link>
			</p>
			<div className="Flex-Row-Space-Between">
				{!object.is_public && (
					<div title="Private">
						<MdLockOutline style={{ cursor: 'pointer', marginRight: '10px', marginTop: '2px' }} />
					</div>
				)}
				{object.is_deleted && (
					<div title="Deleted">
						<MdBlock style={{ cursor: 'pointer', marginRight: '10px', marginTop: '4px' }} />
					</div>
				)}
				<button className="Theme-Btn-Little-One">File</button>
			</div>
		</div>
	);
};

export default ContentBoxItem;
