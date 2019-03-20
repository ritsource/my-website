import React, { useState } from 'react';

import { Data } from '../data/data';

type MyProps = {
	project: Data;
};

const ProjectBoxes = (props: MyProps) => {
	const [ title, setTitle ] = useState();
	const [ emoji, setEmoji ] = useState();
	const [ description, setDescription ] = useState();

	return (
		<div className="ProjectBoxes-c-00">
			<div className="ProjectBoxes-Box-Name-c-01 SearchBox-c-00 Theme-Box-Shadow">
				<h4 className="Flex-Row-Space-Between">
					Project Name
					{false && <div className="Theme-Loading-Spin-Div" />}
				</h4>
				<p className="SearchBox-Description-P-01">Set Title, Emoji and Description for your Project</p>

				<input placeholder="Title" value={title} onChange={(e) => setTitle(e.target.value)} />

				<input placeholder="Emoji" value={emoji} onChange={(e) => setEmoji(e.target.value)} />

				<input placeholder="Description" value={description} onChange={(e) => setDescription(e.target.value)} />
			</div>

			{/* <div className="SearchBox-c-00">
				<h4 className="Flex-Row-Space-Between">
					JSON Data {false && <div className="Theme-Loading-Spin-Div" />}
				</h4>
				<p className="SearchBox-Description-P-01">Quickly modify using JSON data structure.</p>
			</div> */}
		</div>
	);
};

export default ProjectBoxes;
