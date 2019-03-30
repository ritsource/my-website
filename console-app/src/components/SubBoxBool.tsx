import React, { useState, useEffect } from 'react';
import { MdDone } from 'react-icons/md';

import Project from '../types/project';
import Blog from '../types/blog';

type MyProps = {
	object: any; // Project | Blog;
};

const SubBoxName = (props: MyProps) => {
	const { object } = props;

	const [ isPublic, setIsPublic ] = useState(object ? object.is_public : '');
	const [ isDeleted, setIsDeleted ] = useState(object ? object.is_deleted : '');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ boxEditable, setBoxEditable ] = useState(false); // Check if Title, Emoji, or Desc. has changed

	const setPropsValToState = () => {
		setIsPublic(object ? object.is_public : '');
		setIsDeleted(object ? object.is_deleted : '');
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
				Booleans {isAsync && <div className="Theme-Loading-Spin-Div" />}
				{false && <div className="Theme-Loading-Spin-Div" />}
			</h4>
			<p className="SearchBox-Description-P-01">Edit Boolean Values</p>

			<div className="SubBoxDoc-Check-Box-Container-001 Flex-Row-Start">
				<div
					className={`SubBoxDoc-Check-Box ${isPublic && 'SubBoxDoc-Check-Box-Active'}`}
					onClick={() => {
						if (boxEditable) {
							setIsPublic(!isPublic);
						}
					}}
				>
					<MdDone style={{ marginBottom: '-1px' }} />
				</div>
				<p style={{ marginLeft: '10px', padding: '0px' }} className="SearchBox-Description-P-01">
					Is Public
				</p>
			</div>

			<div className="SubBoxDoc-Check-Box-Container-001 Flex-Row-Start">
				<div
					className={`SubBoxDoc-Check-Box ${isDeleted && 'SubBoxDoc-Check-Box-Active'}`}
					onClick={() => {
						if (boxEditable) {
							setIsDeleted(!isDeleted);
						}
					}}
				>
					<MdDone style={{ marginBottom: '-1px' }} />
				</div>
				<p style={{ marginLeft: '10px', padding: '0px' }} className="SearchBox-Description-P-01">
					Is Deleted
				</p>
			</div>

			<div
				style={{
					width: 'calc(100% - 40px)',
					padding: '2px 20px 20px 20px'
				}}
				className="Flex-Row-Start"
			>
				{boxEditable ? (
					<React.Fragment>
						<button className="Theme-Btn-Green">Save</button>
						<button
							style={{ marginLeft: '12px' }}
							className="Theme-Btn-Grey"
							onClick={() => {
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
