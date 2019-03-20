import React, { useState } from 'react';

const SearchBox = () => {
	const [ query, setQuery ] = useState('');

	return (
		<div className="SearchBox-c-00 Theme-Box-Shadow">
			<h4>Search Box</h4>
			<p className="SearchBox-Description-P-01">
				Quick access to all the File-Objects, search on Database. anything from here, all the Projects, Photos,
				Blogs.{' '}
			</p>
			<input placeholder="Search" value={query} onChange={(e) => setQuery(e.target.value)} />
		</div>
	);
};

export default SearchBox;
