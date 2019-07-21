import React, { useState, useEffect } from 'react';
import { MdDone } from 'react-icons/md';

// import Project from '../types/project';
// import Blog from '../types/blog';

type MyProps = {
	object: any; // Project | Blog;
	saveFunction: (c: any, u: any) => void;
	isProject: boolean;
};

const renderInput = (editable: boolean, bool: boolean, setFunc: (b: any) => void, label: string) => {
	return (
		<div className="SubBoxDoc-Check-Box-Container-001 Flex-Row-Start">
			<div
				className={`SubBoxDoc-Check-Box ${bool && 'SubBoxDoc-Check-Box-Active'}`}
				onClick={() => {
					if (editable) {
						setFunc(!bool);
					}
				}}
			>
				<MdDone style={{ marginBottom: '-1px' }} />
			</div>
			<p style={{ marginLeft: '10px', padding: '0px' }} className="SearchBox-Description-P-01">
				{label}
			</p>
		</div>
	);
};

const SubBoxName = (props: MyProps) => {
	const { object, isProject } = props;

	const [ isMajor, setIsMajor ] = useState(object ? object.is_major : '');
	const [ isTechnical, setIsTechnical ] = useState(object ? object.is_technical : '');
	const [ isPublic, setIsPublic ] = useState(object ? object.is_public : '');
	const [ isDeleted, setIsDeleted ] = useState(object ? object.is_deleted : '');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ errorMsg, setErrorMsg ] = useState(false);
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

			{isProject ? (
				renderInput(boxEditable, isMajor, setIsMajor, 'Is Major')
			) : (
				renderInput(boxEditable, isTechnical, setIsTechnical, 'Is Technical')
			)}

			{renderInput(boxEditable, isPublic, setIsPublic, 'Is Public')}
			{renderInput(boxEditable, isDeleted, setIsDeleted, 'Is Deleted')}

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
									// console.log('object', object);
									await props.saveFunction(object, {
										is_major: isMajor,
										is_technical: isTechnical,
										is_public: isPublic,
										is_deleted: isDeleted
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
