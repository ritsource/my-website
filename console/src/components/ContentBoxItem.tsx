import React from 'react';
import { Link } from 'react-router-dom';
import { MdBlock, MdLockOutline, MdLockOpen } from 'react-icons/md';

import { siteBase } from '../api';

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
				<a
					className="a-exception"
					href={
						siteBase +
						'/preview?' +
						`src=${object.doc_type === 'markdown' ? object.markdown : object.html}` +
						'&' +
						`type=${object.doc_type}`
					}
					target="_blank"
				>
					<button className="Theme-Btn-Little-One">Preview</button>
				</a>
			</div>
		</div>
	);
};

export default ContentBoxItem;
