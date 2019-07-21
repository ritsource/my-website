import React, { useState, useEffect } from 'react';
import { MdDone } from 'react-icons/md';

import { serverAddress } from '../api';

type MyProps = {
	object: any; // Project | Blog;
	saveFunction: (c: any, u: any) => void;
};

const SubBoxName = (props: MyProps) => {
	const { object } = props;

	const docTypeOptions = [ 'markdown', 'html' ];
	const [ docType, setDocType ] = useState(object.doc_type);

	const [ html, setHtml ] = useState(object ? object.html : '');
	const [ markdown, setMarkdown ] = useState(object ? object.markdown : '');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ errorMsg, setErrorMsg ] = useState(false);
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
			<p className="SearchBox-Description-P-01">
				Preview the Document{' '}
				<a
					className="a-exception"
					href={
						serverAddress +
						'/preview?' +
						`src=${object.doc_type === 'markdown' ? object.markdown : object.html}` +
						'&' +
						`type=${object.doc_type}`
					}
					target="_blank"
				>
					here
				</a>
			</p>

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
										html,
										markdown,
										doc_type: docType
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
