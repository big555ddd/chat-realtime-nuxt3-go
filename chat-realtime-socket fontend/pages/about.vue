<template>
  <div class="editor-container">
    <!-- Toolbar -->
    <div class="menu-bar" v-if="editor">
      <button @click="toggleBold" :class="{ active: editor.isActive('bold') }"><i class="fas fa-bold"></i></button>
      <button @click="toggleItalic" :class="{ active: editor.isActive('italic') }"><i class="fas fa-italic"></i></button>
      <button @click="toggleUnderline" :class="{ active: editor.isActive('underline') }"><i class="fas fa-underline"></i></button>
      <button @click="toggleStrike" :class="{ active: editor.isActive('strike') }"><i class="fas fa-strikethrough"></i></button>
      <button @click="setParagraph" :class="{ active: editor.isActive('paragraph') }"><i class="fas fa-paragraph"></i></button>
      <button @click="setHeading(1)" :class="{ active: editor.isActive('heading', { level: 1 }) }">H1</button>
      <button @click="setHeading(2)" :class="{ active: editor.isActive('heading', { level: 2 }) }">H2</button>
      <button @click="toggleBulletList" :class="{ active: editor.isActive('bulletList') }"><i class="fas fa-list-ul"></i></button>
      <button @click="toggleOrderedList" :class="{ active: editor.isActive('orderedList') }"><i class="fas fa-list-ol"></i></button>
      <button @click="addLink"><i class="fas fa-link"></i></button>
      <button @click="addImage"><i class="fas fa-image"></i></button>
      <button @click="addCodeBlock" :class="{ active: editor.isActive('codeBlock') }"><i class="fas fa-code"></i></button>
      <button @click="addTable"><i class="fas fa-table"></i></button>
    </div>

    <!-- Editor Content -->
    <EditorContent :editor="editor" class="editor-content" />

    <!-- Save Button -->
    <button @click="saveContent" class="save-button">Save</button>
    <div v-html="savedHTML" class="saved-content"></div>
  </div>
</template>

<script>
import { EditorContent, useEditor } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Bold from '@tiptap/extension-bold'
import Italic from '@tiptap/extension-italic'
import Underline from '@tiptap/extension-underline'
import Strike from '@tiptap/extension-strike'
import Heading from '@tiptap/extension-heading'
import Paragraph from '@tiptap/extension-paragraph'
import BulletList from '@tiptap/extension-bullet-list'
import OrderedList from '@tiptap/extension-ordered-list'
import Link from '@tiptap/extension-link'
import Image from '@tiptap/extension-image'
import CodeBlock from '@tiptap/extension-code-block'
import Table from '@tiptap/extension-table'
import TableRow from '@tiptap/extension-table-row'
import TableCell from '@tiptap/extension-table-cell'
import TableHeader from '@tiptap/extension-table-header'
import '@fortawesome/fontawesome-free/css/all.css'

export default {
  components: {
    EditorContent
  },
  setup() {
    const savedHTML = ref('')
    const editor = useEditor({
      extensions: [
        StarterKit,
        Bold,
        Italic,
        Underline,
        Strike,
        Heading,
        Paragraph,
        BulletList,
        OrderedList,
        Link,
        Image,
        CodeBlock,
        Table.configure({
          resizable: true,
        }),
        TableRow,
        TableCell,
        TableHeader
      ],
      content: '<p>Start editing...</p>',
    })

    const saveContent = () => {
      savedHTML.value = editor.value?.getHTML() || ''
    }

    const toggleBold = () => editor.value?.chain().focus().toggleBold().run()
    const toggleItalic = () => editor.value?.chain().focus().toggleItalic().run()
    const toggleUnderline = () => editor.value?.chain().focus().toggleUnderline().run()
    const toggleStrike = () => editor.value?.chain().focus().toggleStrike().run()
    const setParagraph = () => editor.value?.chain().focus().setParagraph().run()
    const setHeading = level => editor.value?.chain().focus().toggleHeading({ level }).run()
    const toggleBulletList = () => editor.value?.chain().focus().toggleBulletList().run()
    const toggleOrderedList = () => editor.value?.chain().focus().toggleOrderedList().run()
    const addLink = () => {
      const url = prompt('Enter the URL')
      if (url) editor.value?.chain().focus().extendMarkRange('link').setLink({ href: url }).run()
    }
    const addImage = () => {
      const url = prompt('Enter image URL')
      if (url) editor.value?.chain().focus().setImage({ src: url }).run()
    }
    const addCodeBlock = () => editor.value?.chain().focus().toggleCodeBlock().run()
    const addTable = () => editor.value?.chain().focus().insertTable({ rows: 3, cols: 3 }).run()

    return {
      editor,
      saveContent,
      savedHTML,
      toggleBold,
      toggleItalic,
      toggleUnderline,
      toggleStrike,
      setParagraph,
      setHeading,
      toggleBulletList,
      toggleOrderedList,
      addLink,
      addImage,
      addCodeBlock,
      addTable
    }
  }
}
</script>

<style scoped>
.editor-container {
  background-color: #2e2e2e;
  color: white;
  padding: 16px;
  border-radius: 8px;
}
.menu-bar {
  display: flex;
  gap: 8px;
  margin-bottom: 8px;
  background-color: #333;
  padding: 8px;
  border-radius: 4px;
}
.menu-bar button {
  padding: 6px;
  font-weight: bold;
  color: white;
  background-color: transparent;
  border: none;
  cursor: pointer;
}
.menu-bar button.active {
  background-color: #000;
  color: #fff;
  border-radius: 4px;
}
.menu-bar button i {
  font-size: 14px;
}
.editor-content {
  min-height: 300px;
  background-color: #1e1e1e;
  padding: 16px;
  border-radius: 4px;
  color: white;
}
.save-button {
  margin-top: 12px;
  padding: 8px 16px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
.saved-content {
  margin-top: 16px;
  color: white;
}
</style>
