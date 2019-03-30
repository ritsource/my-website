import React, { useState, useEffect } from 'react';

import { Data } from '../data/data';

type MyProps = {
	project: Data;
	pContext: any;
	setProject: () => void;
};

const ProjectBoxes = (props: MyProps) => {
	const { project } = props;
	const [ title, setTitle ] = useState(project ? project.title : '');
	const [ emoji, setEmoji ] = useState(project ? project.emoji : '');
	const [ description, setDescription ] = useState(project ? project.description : '');

	const [ nameBoxEd, setNameBoxEd ] = useState(false); // Check if Title, Emoji, or Desc. has changed
	// const [ nameBoxEd, setNameBoxEd ] = useState();	// Check if Title, Emoji, or Desc. has changed

	const setPropsValToState = () => {
		setTitle(project ? project.title : '');
		setEmoji(project ? project.emoji : '');
		setDescription(project ? project.description : '');
	};

	useEffect(
		() => {
			setPropsValToState();
		},
		[ project ]
	);

	return (
		<div className="ProjectBoxes-c-00">
			<div className="ProjectBoxes-Box-Name-c-01 SearchBox-c-00 Theme-Box-Shadow">
				<h4 className="Flex-Row-Space-Between">
					Project Name
					{false && <div className="Theme-Loading-Spin-Div" />}
				</h4>
				<p className="SearchBox-Description-P-01">Set Title, Emoji and Description for your Project</p>

				<input
					placeholder="Title"
					value={title}
					onChange={(e) => {
						if (nameBoxEd) {
							setTitle(e.target.value);
						}
					}}
				/>

				<input
					placeholder="Emoji"
					value={emoji}
					onChange={(e) => {
						if (nameBoxEd) {
							setEmoji(e.target.value);
						}
					}}
				/>

				<input
					placeholder="Description"
					value={description}
					onChange={(e) => {
						if (nameBoxEd) {
							setDescription(e.target.value);
						}
					}}
				/>

				<div
					style={{
						width: 'calc(100% - 40px)',
						padding: '2px 20px 20px 20px'
					}}
					className="Flex-Row-Start"
				>
					{nameBoxEd ? (
						<React.Fragment>
							<button className="Theme-Btn-Green">Save</button>
							<button
								style={{ marginLeft: '12px' }}
								className="Theme-Btn-Grey"
								onClick={() => {
									setNameBoxEd(false);
									setPropsValToState();
								}}
							>
								Cancel
							</button>
						</React.Fragment>
					) : (
						<button className="Theme-Btn-Main" onClick={() => setNameBoxEd(true)}>
							Edit
						</button>
					)}
				</div>
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
