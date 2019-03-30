import React, { useState, useEffect } from 'react';
import { MdDone } from 'react-icons/md';

import Project from '../types/project';
import Blog from '../types/blog';

type MyProps = {
	object: any; // Project | Blog;
};

const SubBoxName = (props: MyProps) => {
	const { object } = props;

	const docTypeOptions = [ 'markdown', 'html' ];
	const [ docType, setDocType ] = useState(object.doc_type);

	const [ html, setHtml ] = useState(object ? object.html : '');
	const [ markdown, setMarkdown ] = useState(object ? object.markdown : '');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ boxEditable, setBoxEditable ] = useState(false); // Check if Title, Emoji, or Desc. has changed

	const setPropsValToState = () => {
		setHtml(object ? object.html : '');
		setMarkdown(object ? object.markdown : '');
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
				Document {isAsync && <div className="Theme-Loading-Spin-Div" />}
				{false && <div className="Theme-Loading-Spin-Div" />}
			</h4>
			<p className="SearchBox-Description-P-01">Edit Document (Project Documentation or Blog Content)</p>

			<div className="SubBoxDoc-Check-Box-Container-001 Flex-Row-Center">
				<input
					placeholder="Source Link (HTML)"
					value={html}
					onChange={(e) => {
						if (boxEditable) {
							setHtml(e.target.value.trim());
						}
					}}
				/>
				<div
					className={`SubBoxDoc-Check-Box ${docType === 'html' && 'SubBoxDoc-Check-Box-Active'}`}
					onClick={() => {
						if (boxEditable) {
							setDocType('html');
						}
					}}
				>
					<MdDone style={{ marginBottom: '-1px' }} />
				</div>
			</div>

			<div className="SubBoxDoc-Check-Box-Container-001 Flex-Row-Center">
				<input
					placeholder="Source Link (Markdwn)"
					value={markdown}
					onChange={(e) => {
						if (boxEditable) {
							setMarkdown(e.target.value.trim());
						}
					}}
				/>
				<div
					className={`SubBoxDoc-Check-Box ${docType === 'markdown' && 'SubBoxDoc-Check-Box-Active'}`}
					onClick={() => {
						if (boxEditable) {
							setDocType('markdown');
						}
					}}
				>
					<MdDone style={{ marginBottom: '-1px' }} />
				</div>
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
