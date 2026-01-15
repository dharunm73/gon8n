#set page(
  paper: "us-letter",
  margin: (x: 0.75cm, y: 0.75cm),
)

#set text(
  font: "Roboto",
  size: 10pt,
  lang: "en",
  hyphenate: false,
  ligatures: false
)

#set par(justify: false, leading: 0.65em)

#let data = json("updated_resume.json")
#let pdata = json("personal_details.json")

#show heading.where(level: 2): it => {
  v(-0.2cm)
  text(weight: "bold", size: 1.0em, it.body)
  v(-0.35cm)
  line(length: 100%, stroke: 0.5pt + gray)
  v(-0.1cm)
}

#let item-header(title, location, date) = {
  block(width: 100%, below: 0.6em)[
    *#title* #if location != none [ | #location ] 
    #if date != none [ | #date ]
  ]
}

#let highlights(body) = {
  set list(marker: [•], indent: 0pt, body-indent: 0.5em, spacing: 0.5em)
  body
}

#align(center, [
  #text(size: 25pt)[#pdata.name] \
  #v(-5pt)
  #text(size: 10pt)[
    #link("mailto:" + pdata.email)[#pdata.email] |
    #link("tel:" + pdata.entire_phone_number)[#pdata.phone] |
    #pdata.your_location |
    #link(pdata.link_of_linkedin)[LinkedIn: #pdata.linkedin_username] |
    #link(pdata.link_of_github)[GitHub: #pdata.github_username]
  ]
])

#v(-7.3pt)
  #text(size: 14pt, weight: "bold")[Target Role: #data.target_role]
#v(-0.10cm)

== Summary
#eval(data.summary, mode: "markup")

== Education
#item-header(
  pdata.college_name, 
  pdata.college_location, 
  pdata.course_duration
)
#pdata.course_name

== Technical Skills
#for group in data.skills [
  *#group.category:* #group.items \
]

== Relevant Experience & Projects
#for proj in data.projects [
  #let loc = if proj.repo_link != none {
    [Project Repo: #link(proj.repo_link)[GitHub]]
  } else {
    none
  }
  #item-header(
    proj.title, 
    loc,  
    none   
  )

  #v(0.05cm)
  
  #highlights[
    #for point in proj.points [
      - #eval(point, mode: "markup") 
    ]
  ]

  #v(-0.1cm)
]

#v(0.1cm)

== Awards and Achievements
#highlights[
  #for award in data.awards [
    - *#award.title* #if "description" in award and award.description != "" [ (#eval(award.description, mode: "markup")) ], #award.date
  ]
]

== Leadership & Extracurricular Activities
#for role in data.leadership [
  #item-header(
    role.role + " | " + role.org,
    none,
    none
  )
  #highlights[
    #for point in role.points [
      - #eval(point, mode: "markup") 
    ]
  ]
]