

```php
<?php

namespace Tests\Unit;

use App\Helpers\UserBioDataExtractor;
use PHPUnit\Framework\TestCase;

class UserBioDataExtractorTest extends TestCase
{
    public function testUserBioDataExtractorGetName()
    {
        $extractor = new UserBioDataExtractor();
        $bio = $extractor->setName('John Doe');

        $this->assertEquals('John Doe', $bio->getName());
    }

    public function testUserBioDataExtractorGetEmail()
    {
        $extractor = new UserBioDataExtractor();
        $bio = $extractor->setEmail('john.doe@example.com');

        $this->assertEquals('john.doe@example.com', $bio->getEmail());
    }

    public function testUserBioDataExtractorGetAge()
    {
        $extractor = new UserBioDataExtractor();
        $bio = $extractor->setAge(25);

        $this->assertEquals(25, $bio->getAge());
    }

    public function testUserBioDataExtractorGetLocation()
    {
        $extractor = new UserBioDataExtractor();
        $bio = $extractor->setLocation('New York, USA');

        $this->assertEquals('New York, USA', $bio->getLocation());
    }

    public function testUserBioDataExtractorGetBio()
    {
        $extractor = new UserBioDataExtractor();
        $bio = $extractor->setBio('I am a software engineer');

        $this->assertEquals('I am a software engineer', $bio->getBio());
    }
}
```
