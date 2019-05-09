import React, { useState, useEffect } from 'react';

// import Project from '../types/project';
// import Blog from '../types/blog';

const renderInput = (editable: boolean, value: string, setFunc: (b: any) => void, placeholder: string) => {
	return (
		<input
			placeholder={placeholder}
			value={value}
			onChange={(e) => {
				if (editable) {
					setFunc(e.target.value);
				}
			}}
		/>
	);
};

type MyProps = {
	object: any; // Project | Blog;
	saveFunction: (c: any, u: any) => void;
	isProject: boolean; // True if for Project, False for Blogs
};

const SubBoxName = (props: MyProps) => {
	const { object, isProject } = props;

	const [ title, setTitle ] = useState(object ? object.title : '');
	const [ description, setDescription ] = useState(object ? object.description : '');
	const [ description_link, setDescriptionLink ] = useState(object ? object.description_link : '');
	const [ link, setLink ] = useState(object ? object.link : '');
	const [ thumbnail, setThumbnail ] = useState(object ? object.thumbnail : '');
	const [ author, setAuthor ] = useState(object ? object.author : '');
	const [ formatted_date, setFormattedDate ] = useState(object ? object.formatted_date : '');

	const [ isAsync, setIsAsync ] = useState(false); // Is Async
	const [ errorMsg, setErrorMsg ] = useState(false);
	const [ boxEditable, setBoxEditable ] = useState(false); // Check if Title, Emoji, or Desc. has changed

	const setPropsValToState = () => {
		setTitle(object ? object.title : '');
		setDescription(object ? object.description : '');
		setLink(object ? object.link : '');
		setAuthor(object ? object.author : '');
		setFormattedDate(object ? object.formatted_date : '');
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

			{renderInput(boxEditable, title, setTitle, 'Title')}
			{renderInput(boxEditable, description, setDescription, 'Description')}
			{renderInput(boxEditable, description_link, setDescriptionLink, 'A Link in Description (Optional)')}
			{renderInput(boxEditable, thumbnail, setThumbnail, 'Thumbnail Link (For SEO)')}

			{isProject ? (
				renderInput(boxEditable, link, setLink, 'Project Link (Github or App Url)')
			) : (
				<React.Fragment>
					{renderInput(boxEditable, author, setAuthor, 'Author (author name)')}
					{renderInput(boxEditable, formatted_date, setFormattedDate, 'Formatted Release Date')}
				</React.Fragment>
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
										description_link,
										thumbnail,
										link,
										author,
										formatted_date
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
