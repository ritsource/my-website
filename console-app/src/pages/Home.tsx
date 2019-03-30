import React from 'react';

import SearchBox from '../components/SearchBox';
import ContentBox from '../components/ContentBox';

import ProjectContext from '../contexts/ProjectContext';

const HomePage = () => {
	return (
		<ProjectContext.Consumer>
			{(pContext) => {
				return (
					<div className="Page-c-00">
						<div className="Page-Container-00">
							<div style={{ alignItems: 'flex-start' }} className="Flex-Row-Space-Between">
								<SearchBox />
								<ContentBox projects={pContext.projects} readProjects={pContext.readProjects} />
							</div>
						</div>
					</div>
				);
			}}
		</ProjectContext.Consumer>
	);
};

export default HomePage;
