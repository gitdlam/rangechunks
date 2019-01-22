# rangechunks

Sometimes a large list of values is to be searched for and it is impractical to have everything in the SQL statement, but not having a criteria would consume too much unused data.  A compromise is to work out small ranges within the large list, so that the fetched result is small enough, and the SQL statement size is not too great.
