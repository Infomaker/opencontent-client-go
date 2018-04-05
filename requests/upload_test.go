package requests

import (
  "testing"
  "github.com/hansbringert/opencontent-client/ochost"
  "github.com/stretchr/testify/assert"
)

func TestUploadRequest_Setters(t *testing.T) {
  host := ochost.NewOpenContentHost()
  req := NewUploadRequest(host)

  req.Id = "some-id"
  req.Batch = true
  req.Source = "some-source"

  assert.Equal(t, "some-id", req.Id)
  assert.Equal(t, true, req.Batch)
  assert.Equal(t, "some-source", req.Source)

  req.PrimaryFile("/path/to/file/primary-file.xml", "the-primary-file-mimetype")

  // assert the primary file details
  assert.Equal(t, "/path/to/file/primary-file.xml", req.Primary.Fullpath)
  assert.Equal(t, "the-primary-file-mimetype", req.Primary.Mimetype)

  // change to another file assert that the primary file details will change
  req.PrimaryFile("/path/to/file/another-file.xml", "another-mimetype")
  assert.Equal(t, "/path/to/file/another-file.xml", req.Primary.Fullpath)
  assert.Equal(t, "another-mimetype", req.Primary.Mimetype)
  
  // metadata files
  req.MetadataFile("/path/to/file/metadata-file.xml", "the-first-metadata-file-mimetype")
  assert.Equal(t, req.Metadata["metadata-file.xml"].Fullpath, "/path/to/file/metadata-file.xml")
  assert.Equal(t, req.Metadata["metadata-file.xml"].Mimetype, "the-first-metadata-file-mimetype")

  req.MetadataFile("/path/to/file/metadata-2-file.xml", "the-second-metadata-file-mimetype")
  assert.Equal(t, req.Metadata["metadata-2-file.xml"].Fullpath, "/path/to/file/metadata-2-file.xml")
  assert.Equal(t, req.Metadata["metadata-2-file.xml"].Mimetype, "the-second-metadata-file-mimetype")

  req.MetadataFile("/path/to/file/metadata-3-file.xml", "the-third-metadata-file-mimetype")
  assert.Equal(t, req.Metadata["metadata-3-file.xml"].Fullpath, "/path/to/file/metadata-3-file.xml")
  assert.Equal(t, req.Metadata["metadata-3-file.xml"].Mimetype, "the-third-metadata-file-mimetype")

  // create the files map
  filesmap := req.createFilesMap()
  // assert files map
  assert.Equal(t, "/path/to/file/another-file.xml", filesmap["another-file.xml"])
  assert.Equal(t, "/path/to/file/metadata-file.xml", filesmap["metadata-file.xml"])
  assert.Equal(t, "/path/to/file/metadata-2-file.xml", filesmap["metadata-2-file.xml"])
  assert.Equal(t, "/path/to/file/metadata-3-file.xml", filesmap["metadata-3-file.xml"])
  assert.Equal(t, 4, len(filesmap))

  // create the form map
  formmap := req.createFormMap()

  // assert the form map
  assert.Equal(t, "another-mimetype", formmap["file-mimetype"])
  assert.Equal(t, "metadata-file.xml", formmap["metadata"])
  assert.Equal(t, "the-first-metadata-file-mimetype", formmap["metadata-mimetype"])

  assert.Equal(t, "metadata-2-file.xml", formmap["metadata2"])
  assert.Equal(t, "the-second-metadata-file-mimetype", formmap["metadata2-mimetype"])

  assert.Equal(t, "metadata-3-file.xml", formmap["metadata3"])
  assert.Equal(t, "the-third-metadata-file-mimetype", formmap["metadata3-mimetype"])

}
