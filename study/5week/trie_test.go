package trie

run test|debug test
func TestInsert(t *testing.T){
	root := &NewNode{""}
	success := Insert(root, "tea")
	assert.True(t, success)

	success := Insert(root, "ted")
	assert.True(t, success)

	success := Insert(root, "ten")
	assert.True(t, success)

	err := drawer.SaveTreeGraph(root, "./tree.png")
	assert.Nil(t,err)
}