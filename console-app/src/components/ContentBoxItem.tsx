import React from 'react';

import { Data } from './ContentBox';

type MyProps = {
	object: Data;
};

const ContentBoxItem: React.SFC<MyProps> = (props) => {
	const { title, emoji, description, html, markdown, link, imageUrl } = props.object;

	return (
		<div className="ContentBoxItem-c-00 Flex-Row-Space-Between">
			<p>
				<span>{emoji}</span>
				<a href={link} target="_blank">
					{title}
				</a>
			</p>
			<button className="Theme-Btn-Little-One">File</button>
		</div>
	);
};

export default ContentBoxItem;
