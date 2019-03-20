import React, { useState } from 'react';

const SearchBox = () => {
	const [ query, setQuery ] = useState('');

	return (
		<div className="SearchBox-c-00 Theme-Box-Shadow Flex-Column-Start">
			<h4>Search Box</h4>
			<p>You can search anything from here.. Projects, Photos, Blogs..</p>
			<input placeholder="Search" value={query} onChange={(e) => setQuery(e.target.value)} />
		</div>
	);
};

export default SearchBox;
