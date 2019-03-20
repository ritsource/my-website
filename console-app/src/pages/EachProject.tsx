import React, { useState, useEffect } from 'react';

import JSONBox from '../components/JSONBox';
import ProjectBoxes from '../components/ProjectBox';
import data, { Data } from '../data/data';

type MyProps = {
	history: any;
	match: any;
};

const EachProjectPage = (props: MyProps) => {
	const [ project, setProject ]: any = useState(null);

	useEffect(() => {
		setProject(data.find(({ _id }) => _id === props.match.params.projectId));
		console.log(data.find(({ _id }) => _id === props.match.params.projectId));
	}, []);

	useEffect(
		() => {
			console.log(project);
		},
		[ project ]
	);

	return (
		<div className="Page-c-00">
			<div className="Page-Container-00">
				<h1>Project / {props.match.params.projectId}</h1>
				<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
					<JSONBox object={project} />
					<div className="Page-Vertical-Box-Container">
						<ProjectBoxes project={project} />
					</div>
				</div>
			</div>
		</div>
	);
};

export default EachProjectPage;
