import React, { useState, useEffect } from 'react';

import Project from '../types/project';
import Blog from '../types/blog';

type MyProps = {
	object: any; // Project | Blog;
	saveFunction: (c: any, u: any) => void;
};

const SubBoxName = (props: MyProps) => {
	const { object } = props;

	const [ title, setTitle ] = useState(object ? object.title : '');
	const [ description, setDescription ] = useState(object ? object.description : '');
	const [ link, setLink ] = useState(object ? object.link : '');
	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ errorMsg, setErrorMsg ] = useState(false);
	const [ boxEditable, setBoxEditable ] = useState(false); // Check if Title, Emoji, or Desc. has changed

	const setPropsValToState = () => {
		setTitle(object ? object.title : '');
		setLink(object ? object.link : '');
		setDescription(object ? object.description : '');
	};

	useEffect(
		() => {
			setPropsValToState();
		},
		[ object ]
	);

	return (
		<div className="SubBoxName-c-00 SearchBox-c-00 Theme-Box-Shadow">
			<h4 className="Flex-Row-Space-Between">
				{title ? title : 'Project Name, Description'} {isAsync && <div className="Theme-Loading-Spin-Div" />}
			</h4>
			<p className="SearchBox-Description-P-01">Set Title, Emoji and Description for your Project</p>

			<input
				placeholder="Title"
				value={title}
				onChange={(e) => {
					if (boxEditable) {
						setTitle(e.target.value);
					}
				}}
			/>

			<input
				placeholder="Description"
				value={description}
				onChange={(e) => {
					if (boxEditable) {
						setDescription(e.target.value);
					}
				}}
			/>

			{(!!link || link === '') && (
				<input
					placeholder="Project Link (Github or App Url)"
					value={link}
					onChange={(e) => {
						if (boxEditable) {
							setLink(e.target.value.trim());
						}
					}}
				/>
			)}

			{errorMsg && (
				<p
					style={{
						color: 'var(--danger-red-color)',
						padding: '0px 20px 10px 20px'
					}}
					className="SearchBox-Description-P-01"
				>
					Error: {errorMsg}
				</p>
			)}

			<div
				style={{
					width: 'calc(100% - 40px)',
					padding: '2px 20px 20px 20px'
				}}
				className="Flex-Row-Start"
			>
				{boxEditable ? (
					<React.Fragment>
						<button
							className="Theme-Btn-Green"
							onClick={async () => {
								setIsAsync(true);
								try {
									await props.saveFunction(object, {
										title,
										description,
										link
									});
									setErrorMsg(false);
									setBoxEditable(false);
								} catch (e) {
									setErrorMsg(e);
								}
								setIsAsync(false);
							}}
						>
							Save
						</button>
						<button
							style={{ marginLeft: '12px' }}
							className="Theme-Btn-Grey"
							onClick={() => {
								setErrorMsg(false);
								setBoxEditable(false);
								setPropsValToState();
							}}
						>
							Cancel
						</button>
					</React.Fragment>
				) : (
					<button className="Theme-Btn-Main" onClick={() => setBoxEditable(true)}>
						Edit
					</button>
				)}
			</div>
		</div>
	);
};

export default SubBoxName;
